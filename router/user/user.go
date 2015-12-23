package user

import (
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
)

type userRouter struct {
	routes []router.Route
}

func NewRouter() router.Router {
	r := &userRouter{}

	r.initRoutes()

	return r
}

func (r *userRouter) Routes() []router.Route {
	return r.routes
}

func (r *userRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute(
			"user/findAll/",
			r.FindAllAvailable,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewGetRoute(
			"user/findById/{id:[0-9]+}",
			r.FindById,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewGetRoute(
			"user/findByUsername/{username:[0-9a-zA-Z]+}",
			r.FindByUsername,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewPostRoute(
			"user/edit/",
			r.Edit,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewPostRoute(
			"user/create/",
			r.Create,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewPostRoute(
			"user/changePassword/",
			r.ChangePassword,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewPostRoute(
			"user/delete/",
			r.Delete,
			true,
			[]string{"ADMIN"},
			httputils.API,
		),
		router.NewGetRoute(
			"user/profile/",
			r.GetProfileForCurrentUser,
			true,
			[]string{},
			httputils.API,
		),
		router.NewPostRoute(
			"user/current/changePassword/",
			r.ChangePasswordForCurrentUser,
			true,
			[]string{},
			httputils.API,
		),
		router.NewPostRoute(
			"user/current/changeEmail/",
			r.ChangeEmailForCurrentUser,
			true,
			[]string{},
			httputils.API,
		),
		router.NewPostRoute(
			"user/current/changeFullName/",
			r.ChangeFullNameForCurrentUser,
			true,
			[]string{},
			httputils.API,
		),
		router.NewGetRoute(
			"user/current/privileges/",
			r.GetPrivilegesForCurrentUser,
			true,
			[]string{},
			httputils.API,
		),
	}
}
