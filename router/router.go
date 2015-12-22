package router

import "github.com/ab22/abcd/router/httputils"

type Router interface {
	Routes() []Route
}

type Route interface {
	Path() string
	Method() string
	Handler() httputils.APIHandler
	RequiresAuth() bool
	RequiredRoles() []string
	Type() httputils.RouteType
}
