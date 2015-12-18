package main

import (
	"net/http"
	"path"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/handlers"
	"github.com/gorilla/mux"
)

type Router struct {
	muxRouter *mux.Router
}

// Registers a route to a template file. Sets up authentication middleware
// to validate session if it is needed.
func (r *Router) registerTemplateRoutes() {
	for _, route := range routes.TemplateRoutes {
		handlerFunc := handlers.GzipContent(route.HandlerFunc)

		if route.RequiresAuth {
			handlerFunc = handlers.ValidateAuth(handlerFunc)
		}

		r.muxRouter.
			Methods(route.Method).
			Path(route.Pattern).
			HandlerFunc(handlerFunc)
	}
}

// Registers API routes. Basically, it just makes a call to
// handlers.JsonHandler to process the handlers response.
func (r *Router) registerApiRoutes() {
	for _, route := range routes.ApiRoutes {
		handlerFunc := handlers.JsonHandler(route.HandlerFunc)
		handlerFunc = handlers.GzipContent(handlerFunc)

		if route.RequiresAuth {
			if len(route.RequiredRoles) > 0 {
				handlerFunc = handlers.Authorize(route.RequiredRoles, handlerFunc)
			}

			handlerFunc = handlers.ValidateAuth(handlerFunc)
		}

		r.muxRouter.
			Methods(route.Method).
			Path("/api/" + route.Pattern).
			HandlerFunc(handlerFunc)
	}
}

// registerStaticFilesServer creates the static file server for the default
// admin app. This might need to change in case we add more frontend apps.
func (r *Router) registerStaticFilesServer() {
	// Register the static files server handler separately.
	var (
		adminAppPath       = config.EnvVariables.App.Frontend.Admin
		staticFilesPath    = path.Join(adminAppPath, "static")
		staticFilesHandler = handlers.NoDirListing(http.FileServer(http.Dir(staticFilesPath)))
	)

	staticFilesHandler = handlers.GzipContent(staticFilesHandler)

	r.muxRouter.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", staticFilesHandler))
}

// Directly call the mux muxRouter's ServeHttp function since our muxRouter is
// just masking mux's muxRouter. We create a ServeHTTP function so that our
// Router becomes a http.Handler and can be passed in as a parameter
// to the http.ListenAndServe function.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.muxRouter.ServeHTTP(w, req)
}

// Creates a new Router and registers all routes to their handlers
// and includes all necessary middleware functions to each of the routes.
func NewRouter() *Router {
	r := &Router{
		muxRouter: mux.NewRouter().StrictSlash(true),
	}

	r.registerTemplateRoutes()
	r.registerApiRoutes()
	r.registerStaticFilesServer()

	return r
}
