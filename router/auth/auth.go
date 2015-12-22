package auth

import (
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
)

type authRouter struct {
	routes []router.Route
}

func NewRouter() router.Router {
	r := &authRouter{}

	r.initRoutes()

	return r
}

func (r *authRouter) Routes() []router.Route {
	return r.routes
}

func (r *authRouter) initRoutes() {
	r.routes = []router.Route{
		httputils.NewGetRoute(
			"auth/checkAuthentication/",
			r.CheckAuth,
			true,
			[]string{},
		),
		httputils.NewPostRoute(
			"auth/login/",
			r.Login,
			false,
			[]string{},
		),
		httputils.NewPostRoute(
			"auth/logout",
			r.Logout,
			true,
			[]string{},
		),
	}
}
