package main

import (
	"net/http"

	"github.com/ab22/abcd/handlers"
)

// Defines an API route with an ApiHandler type as HandlerFunc instead
// of a http.HandleFunc.
type ApiRoute struct {
	Pattern      string
	Method       string
	HandlerFunc  handlers.ApiHandler
	RequiresAuth bool
}

// Simple Template route structure.
type TemplateRoute struct {
	Pattern      string
	Method       string
	HandlerFunc  http.HandlerFunc
	RequiresAuth bool
}

// Routes contains a list of template routes and a list of api routes
// to be registered with middleware as needed.
// Template routes are just routes that lead to handlers that serve html
// templates or static files.
// Api Routes receive and respond json.
type Routes struct {
	TemplateRoutes []TemplateRoute
	ApiRoutes      []ApiRoute
}

var routes = Routes{
	TemplateRoutes: []TemplateRoute{
		{
			Pattern:      "/",
			Method:       "GET",
			HandlerFunc:  handlers.StaticHandler.IndexHandler,
			RequiresAuth: false,
		},
	},
	ApiRoutes: []ApiRoute{},
}
