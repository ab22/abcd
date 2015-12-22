package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/auth"
	"github.com/ab22/abcd/router/httputils"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

type Server struct {
	port            int
	muxRouter       *mux.Router
	bootstrapper    bootstrapper
	cookieStore     *sessions.CookieStore
	cachedTemplates *template.Template
	routers         []router.Router
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
	s.cacheTemplates()

	s.port = config.EnvVariables.App.Port

	log.Println("Setting up routes...")
	s.configureRouters()
	s.createMuxRouter()

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

func (s *Server) cacheTemplates() {
	adminAppPath := config.EnvVariables.App.Frontend.Admin

	s.cachedTemplates = template.Must(template.ParseFiles(adminAppPath + "/index.html"))
}

func (s *Server) addRouter(r router.Router) {
	s.routers = append(s.routers, r)
}

func (s *Server) configureRouters() {
	s.addRouter(auth.NewRouter())
}

func (s *Server) createMuxRouter() {
	s.muxRouter = mux.NewRouter().StrictSlash(true)

	for _, apiRoute := range s.routers {
		for _, r := range apiRoute.Routes() {
			httpHandler := s.makeHttpHandler(r)

			s.muxRouter.
				Methods(r.Method()).
				Path("/api/" + r.Path()).
				HandlerFunc(httpHandler)
		}
	}
}

func (s *Server) makeHttpHandler(route router.Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		handlerFunc := s.handleWithGlobalMiddlewares(route)
		err := handlerFunc(ctx, w, r)

		if err != nil {
			log.Printf("Handler [%s][%s] returned error: %s", r.Method, r.URL.Path, err)
		}
	}
}

func (s *Server) handleWithGlobalMiddlewares(route router.Route) httputils.APIHandler {
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
