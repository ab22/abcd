package student

import (
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
)

type studentRouter struct {
	routes []router.Route
}

// NewRouter creates a new router for the studentRouter.
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
		router.NewPostRoute(
			"student/edit/",
			r.Edit,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewPostRoute(
			"student/create/",
			r.Create,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewGetRoute(
			"student/findByIdNumber/{idNumber:[0-9A-Za-z-]+}",
			r.FindByIdNumber,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
	}
}
