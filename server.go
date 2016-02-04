package main

import (
	// Standard libs
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path"

	// Routes
	"github.com/ab22/abcd/router/auth"
	"github.com/ab22/abcd/router/static"
	"github.com/ab22/abcd/router/student"
	"github.com/ab22/abcd/router/user"

	// Misc.
	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

// Server contains instance details for the server.
type Server struct {
	cookieStore *sessions.CookieStore
	muxRouter   *mux.Router
	services    *services.Services
	routers     []router.Router
	cfg         *config.Config
}

// NewServer returns a new instance of the server. All server configuration
// is done at the Configure() function.
func NewServer() *Server {
	return &Server{}
}

// Configure initializes all necessary data for the server, including the
// configuration data, services and routes.
func (s *Server) Configure() error {
	var err error

	log.Println("Configuring server...")

	s.cfg = config.NewConfig()
	s.cfg.Print()

	err = s.cfg.Validate()

	if err != nil {
		return err
	}

	err = s.configureServices()

	if err != nil {
		return err
	}

	s.configureCookieStore()

	log.Println("Setting up routes...")
	s.createMuxRouter()

	s.configureRouters()
	s.bindRoutes()

	log.Println("Creating static file server...")
	s.createStaticFilesServer()

	return nil
}

// ListenAndServe attaches the current server to the specified configuration
// port.
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%d", s.cfg.App.Port),
		s.muxRouter,
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

	gob.Register(&router.SessionData{})

	s.cookieStore = sessions.NewCookieStore([]byte(secretKey))
	s.cookieStore.MaxAge(0)
}

// addRouter appends a router to the server's router.
func (s *Server) addRouter(r router.Router) {
	s.routers = append(s.routers, r)
}

// configureRouters adds all routers to the server.
func (s *Server) configureRouters() {
	var appPath = s.cfg.App.Frontend.Admin

	s.addRouter(static.NewRouter(appPath))
	s.addRouter(auth.NewRouter())
	s.addRouter(user.NewRouter())
	s.addRouter(student.NewRouter())
}

// createMuxRouter initializes the server's router.
func (s *Server) createMuxRouter() {
	s.muxRouter = mux.NewRouter().StrictSlash(true)
}

// bindRoutes adds all routes to the server's router.
func (s *Server) bindRoutes() {
	for _, route := range s.routers {
		for _, r := range route.Routes() {
			var routePath string
			httpHandler := s.makeHttpHandler(r)

			if r.Type() == httputils.API {
				routePath = "/api/" + r.Path()
			} else {
				routePath = r.Path()
			}

			s.muxRouter.
				Methods(r.Method()).
				Path(routePath).
				HandlerFunc(httpHandler)
		}
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

	contextHandler = router.NoDirListing(router.GzipContent(contextHandler))

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.Background()
		err := contextHandler(ctx, w, req)

		if err != nil {
			log.Printf("static file handler [%s][%s] returned error: %s", req.Method, req.URL.Path, err)
			httputils.WriteError(w, http.StatusInternalServerError, "")
		}
	})

	s.muxRouter.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", handler))
}

// makeHttpHandler creates a http.HandlerFunc from a httputils.ContextHandler.
func (s *Server) makeHttpHandler(route router.Route) http.HandlerFunc {
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
func (s *Server) handleWithMiddlewares(route router.Route) httputils.ContextHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		serverCtx := context.WithValue(ctx, "cookieStore", s.cookieStore)
		serverCtx = context.WithValue(serverCtx, "route", route)
		serverCtx = context.WithValue(serverCtx, "services", s.services)
		serverCtx = context.WithValue(serverCtx, "config", s.cfg)

		h := route.Handler()
		h = router.HandleHttpError(h)
		h = router.GzipContent(h)

		if route.RequiresAuth() {
			if requiredRoles := route.RequiredRoles(); len(requiredRoles) > 0 {
				h = router.Authorize(h)
			}

			h = router.ValidateAuth(h)
		}

		return h(serverCtx, w, r)
	}
}
