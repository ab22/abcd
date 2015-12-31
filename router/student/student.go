package student

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

func (r *studentRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute(
			"student/findAll",
			r.FindAllAvailable,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewGetRoute(
			"student/findById/{id:[0-9]+}",
			r.FindById,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
	}
}
