package main

import (
	// Standard libs
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/handlers"
	"github.com/ab22/abcd/httputils"
	"github.com/ab22/abcd/routes"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

// Server contains instance details for the server.
type Server struct {
	cfg         *config.Config
	cookieStore *sessions.CookieStore
	router      *mux.Router
	services    *services.Services
}

// NewServer returns a new instance of the server. All server configuration
// is done at the Configure() function.
func NewServer() *Server {
	var (
		err    error
		server = &Server{}
	)

	log.Println("Configuring server...")

	server.cfg = config.NewConfig()
	server.cfg.Print()

	err = server.cfg.Validate()

	if err != nil {
		log.Fatalln(err)
	}

	err = server.configureServices()

	if err != nil {
		log.Fatalln(err)
	}

	server.configureCookieStore()

	log.Println("Setting up routes...")

	server.configureRouter()

	log.Println("Creating static file server...")
	server.createStaticFilesServer()

	return server
}

// ListenAndServe attaches the current server to the specified configuration
// port.
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%d", s.cfg.App.Port),
		s.router,
	)
}

// createDatabaseConn creates a new GORM database with the specified database
// configuration.
func (s *Server) createDatabaseConn() (*gorm.DB, error) {
	var (
		db               gorm.DB
		err              error
		dbCfg            = s.cfg.Db
		connectionString = fmt.Sprintf(
			"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.User,
			dbCfg.Password,
			dbCfg.Name,
		)
	)

	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(s.cfg.DbLogMode)

	return &db, nil
}

// configureServices creates the new services for the application to use.
func (s *Server) configureServices() error {
	db, err := s.createDatabaseConn()

	if err != nil {
		return err
	}

	s.services = services.NewServices(db)
	return nil
}

// configureCookieStore creates the cookie store used to validate user
// sessions.
func (s *Server) configureCookieStore() {
	secretKey := s.cfg.App.Secret

	gob.Register(&httputils.SessionData{})

	s.cookieStore = sessions.NewCookieStore([]byte(secretKey))
	s.cookieStore.MaxAge(0)
}

// configureRouter initializes the server's router.
func (s *Server) configureRouter() {
	s.router = mux.NewRouter().StrictSlash(true)

	r := routes.NewRoutes(s.cfg, s.services)

	s.bindRoutes(r.TemplateRoutes, false)
	s.bindRoutes(r.APIRoutes, true)
}

// bindRoutes adds all routes to the server's router.
func (s *Server) bindRoutes(r []routes.Route, apiRoute bool) {
	for _, route := range r {
		var routePath string
		httpHandler := s.makeHTTPHandler(route)

		if apiRoute {
			routePath = "/api/" + route.Pattern()
		} else {
			routePath = route.Pattern()
		}

		s.router.
			Methods(route.Method()).
			Path(routePath).
			HandlerFunc(httpHandler)
	}
}

// createStaticFilesServer creates a static file server to server all of the
// frontend files (html, js, css, etc).
func (s *Server) createStaticFilesServer() {
	var (
		adminAppPath    = s.cfg.App.Frontend.Admin
		staticFilesPath = path.Join(adminAppPath, "static")
	)

	contextHandler := func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		file := path.Join(staticFilesPath, req.URL.Path)

		http.ServeFile(w, req, file)
		return nil
	}

	contextHandler = handlers.HandleHTTPError(contextHandler)
	contextHandler = handlers.GzipContent(contextHandler)
	contextHandler = handlers.NoDirListing(contextHandler)

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.Background()
		err := contextHandler(ctx, w, req)

		if err != nil {
			log.Printf("static file handler [%s][%s] returned error: %s", req.Method, req.URL.Path, err)
		}
	})

	s.router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", handler))
}

// makeHTTPHandler creates a http.HandlerFunc from a httputils.ContextHandler.
func (s *Server) makeHTTPHandler(route routes.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err         error
			handlerFunc httputils.ContextHandler

			ctx = context.Background()
		)

		handlerFunc = s.handleWithMiddlewares(route)
		err = handlerFunc(ctx, w, r)

		if err != nil {
			log.Printf("Handler [%s][%s] returned error: %s", r.Method, r.URL.Path, err)
		}
	}
}

// handleWithMiddlewares applies all middlewares to the specified route. Some
// middleware functions are applied depending on the route's properties, such
// as ValidateAuth and Authorize middlewares. These last 2 functions require
// that the route RequiresAuth() and that RequiredRoles() > 0.
func (s *Server) handleWithMiddlewares(route routes.Route) httputils.ContextHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		serverCtx := context.WithValue(ctx, "cookieStore", s.cookieStore)
		serverCtx = context.WithValue(serverCtx, "route", route)
		serverCtx = context.WithValue(serverCtx, "config", s.cfg)

		h := route.HandlerFunc()
		h = handlers.HandleHTTPError(h)
		h = handlers.GzipContent(h)

		if route.RequiresAuth() {
			if requiredRoles := route.RequiredRoles(); len(requiredRoles) > 0 {
				h = handlers.Authorize(h)
			}

			h = handlers.ValidateAuth(h)
		}

		return h(serverCtx, w, r)
	}
}
