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
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

type Server struct {
	port         int
	muxRouter    *mux.Router
	bootstrapper bootstrapper
	cookieStore  *sessions.CookieStore
	routers      []router.Router
}

func NewServer() *Server {
	return &Server{
		bootstrapper: NewBootstrapper(),
	}
}

func (s *Server) Configure() error {
	log.Println("Configuring server...")

	err := s.bootstrapper.Configure()
	if err != nil {
		return err
	}

	s.configureCookieStore()

	s.port = config.EnvVariables.App.Port

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
		fmt.Sprintf(":%d", s.port),
		s.muxRouter,
	)
}

func (s *Server) configureCookieStore() {
	secretKey := config.EnvVariables.App.Secret

	gob.Register(&router.SessionData{})

	s.cookieStore = sessions.NewCookieStore([]byte(secretKey))
	s.cookieStore.MaxAge(30 * 60)
}

func (s *Server) addRouter(r router.Router) {
	s.routers = append(s.routers, r)
}

func (s *Server) configureRouters() {
	var appPath = config.EnvVariables.App.Frontend.Admin

	s.addRouter(static.NewRouter(appPath))
	s.addRouter(auth.NewRouter())
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

			log.Printf("Binding route [%s][%s]\n", r.Method(), routePath)

			s.muxRouter.
				Methods(r.Method()).
				Path(routePath).
				HandlerFunc(httpHandler)
		}
	}
}

func (s *Server) createStaticFilesServer() {
	var (
		adminAppPath    = config.EnvVariables.App.Frontend.Admin
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

		if route.Type() == httputils.API {
			handlerFunc = s.handleAPIWithGlobalMiddlewares(route)
		} else {
			handlerFunc = s.handleStaticWithGlobalMiddleWares(route)
		}

		err = handlerFunc(ctx, w, r)

		if err != nil {
			log.Printf("Handler [%s][%s] returned error: %s", r.Method, r.URL.Path, err)
		}
	}
}

func (s *Server) handleStaticWithGlobalMiddleWares(route router.Route) httputils.APIHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		middlewares := []router.MiddlewareFunc{
			router.GzipContent,
		}

		serverCtx := context.WithValue(ctx, "cookieStore", s.cookieStore)
		serverCtx = context.WithValue(serverCtx, "route", route)

		h := route.Handler()

		for _, middlewareFunc := range middlewares {
			h = middlewareFunc(h)
		}

		return h(serverCtx, w, r)
	}
}

func (s *Server) handleAPIWithGlobalMiddlewares(route router.Route) httputils.APIHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		middlewares := []router.MiddlewareFunc{
			router.GzipContent,
			router.ValidateAuth,
		}

		serverCtx := context.WithValue(ctx, "cookieStore", s.cookieStore)
		serverCtx = context.WithValue(serverCtx, "route", route)

		h := route.Handler()

		for _, middlewareFunc := range middlewares {
			h = middlewareFunc(h)
		}

		return h(serverCtx, w, r)
	}
}
