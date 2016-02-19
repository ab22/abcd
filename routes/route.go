package routes

import (
	"github.com/ab22/abcd/handlers/httputils"
)

// Route interface.
type Route interface {
	Pattern() string
	Method() string
	HandlerFunc() httputils.ContextHandler
	RequiresAuth() bool
}

type route struct {
	pattern      string
	method       string
	handlerFunc  httputils.ContextHandler
	requiresAuth bool
}

func (r *route) Pattern() string {
	return r.pattern
}

func (r *route) Method() string {
	return r.method
}

func (r *route) HandlerFunc() httputils.ContextHandler {
	return r.handlerFunc
}

func (r *route) RequiresAuth() bool {
	return r.requiresAuth
}
