package main

import (
	"net/http"
	"path"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/handlers"
	"github.com/gorilla/mux"
)

// Registers a route to a template file. Sets up authentication middleware
// to validate session if it is needed.
func registerTemplateRoutes(router *mux.Router) {
	for _, route := range routes.TemplateRoutes {
		handlerFunc := handlers.GzipContent(route.HandlerFunc)

		if route.RequiresAuth {
			handlerFunc = handlers.ValidateAuth(handlerFunc)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			HandlerFunc(handlerFunc)
	}
}

// Registers API routes. Basically, it just makes a call to
// handlers.JsonHandler to process the handlers response.
func registerApiRoutes(router *mux.Router) {
	for _, route := range routes.ApiRoutes {
		handlerFunc := handlers.JsonHandler(route.HandlerFunc)
		handlerFunc = handlers.GzipContent(handlerFunc)

		if route.RequiresAuth {
			handlerFunc = handlers.ValidateAuth(handlerFunc)
		}

		router.
			Methods(route.Method).
			Path("/api/" + route.Pattern).
			HandlerFunc(handlerFunc)
	}
}

func registerStaticFilesServer(router *mux.Router) {
	// Register the static files server handler separately.
	adminPath := config.EnvVariables.App.Frontend.Admin
	staticFilesPath := path.Join(adminPath, "static")

	staticFilesHandler := handlers.NoDirListing(http.FileServer(http.Dir(staticFilesPath)))
	staticFilesHandler = handlers.GzipContent(staticFilesHandler)

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", staticFilesHandler))
}

// Registers all routes to their handlers and makes sure to call
// all necessary middleware functions such as noDirListing.
func registerRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	registerTemplateRoutes(router)
	registerApiRoutes(router)
	registerStaticFilesServer(router)

	return router
}
