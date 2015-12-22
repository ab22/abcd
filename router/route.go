package router

import (
	"github.com/ab22/abcd/router/httputils"
)

type Route interface {
	Path() string
	Method() string
	Handler() httputils.ContextHandler
	RequiresAuth() bool
	RequiredRoles() []string
	Type() httputils.RouteType
}

type route struct {
	method        string
	path          string
	handler       httputils.ContextHandler
	requiresAuth  bool
	requiredRoles []string
	routeType     httputils.RouteType
}

func (r *route) Path() string {
	return r.path
}

func (r *route) Method() string {
	return r.method
}

func (r *route) Handler() httputils.ContextHandler {
	return r.handler
}

func (r *route) RequiresAuth() bool {
	return r.requiresAuth
}

func (r *route) RequiredRoles() []string {
	return r.requiredRoles
}

func (r *route) Type() httputils.RouteType {
	return r.routeType
}

func NewRoute(method, path string, handler httputils.ContextHandler, requiresAuth bool, requiredRoles []string, routeType httputils.RouteType) *route {
	return &route{
		method:        method,
		path:          path,
		handler:       handler,
		requiresAuth:  requiresAuth,
		requiredRoles: requiredRoles,
		routeType:     routeType,
	}
}

func NewGetRoute(path string, handler httputils.ContextHandler, requiresAuth bool, requiredRoles []string, routeType httputils.RouteType) *route {
	return NewRoute("GET", path, handler, requiresAuth, requiredRoles, routeType)
}

func NewPostRoute(path string, handler httputils.ContextHandler, requiresAuth bool, requiredRoles []string, routeType httputils.RouteType) *route {
	return NewRoute("POST", path, handler, requiresAuth, requiredRoles, routeType)
}
