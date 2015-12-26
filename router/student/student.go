package auth

import (
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
)

type studentRouter struct {
	routes []router.Route
}

func NewRouter() router.Router {
	r := &studentRouter{}

	r.initRoutes()

	return r
}

func (r *studentRouter) Routes() []router.Route {
	return r.routes
}

func (r *authRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute(
			"student/findAll",
			r.FindAllAvailable,
			false,
			[]string{},
			httputils.API,
		),
	}
}
