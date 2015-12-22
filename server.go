package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/auth"
	"github.com/ab22/abcd/router/httputils"
	"github.com/ab22/abcd/router/static"
	"github.com/ab22/abcd/router/user"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

type Server struct {
	cookieStore *sessions.CookieStore
	muxRouter   *mux.Router
	services    *services.Services
	routers     []router.Router
	cfg         *config.Config
}

func NewServer() *Server {
	return &Server{}
}

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

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%d", s.cfg.App.Port),
		s.muxRouter,
	)
}

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

func (s *Server) configureServices() error {
	db, err := s.createDatabaseConn()

	if err != nil {
		return err
	}

	s.services = services.NewServices(db)
	return nil
}

func (s *Server) configureCookieStore() {
	secretKey := s.cfg.App.Secret

	gob.Register(&router.SessionData{})

	s.cookieStore = sessions.NewCookieStore([]byte(secretKey))
	s.cookieStore.MaxAge(30 * 60)
}

func (s *Server) addRouter(r router.Router) {
	s.routers = append(s.routers, r)
}

func (s *Server) configureRouters() {
	var appPath = s.cfg.App.Frontend.Admin

	s.addRouter(static.NewRouter(appPath))
	s.addRouter(auth.NewRouter())
	s.addRouter(user.NewRouter())
}

func (s *Server) createMuxRouter() {
	s.muxRouter = mux.NewRouter().StrictSlash(true)
}

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
		}
	})

	s.muxRouter.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", handler))
}

func (s *Server) makeHttpHandler(route router.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err         error
			handlerFunc httputils.APIHandler

			ctx = context.Background()
		)

		handlerFunc = s.handleWithMiddlewares(route)
		err = handlerFunc(ctx, w, r)

		if err != nil {
			log.Printf("Handler [%s][%s] returned error: %s", r.Method, r.URL.Path, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (s *Server) handleWithMiddlewares(route router.Route) httputils.APIHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		serverCtx := context.WithValue(ctx, "cookieStore", s.cookieStore)
		serverCtx = context.WithValue(serverCtx, "route", route)
		serverCtx = context.WithValue(serverCtx, "services", s.services)

		h := route.Handler()
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
