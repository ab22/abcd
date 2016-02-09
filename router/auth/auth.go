package auth

import (
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
)

type authRouter struct {
	routes []router.Route
}

// NewRouter creates a new router for the authRouter.
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
		router.NewPostRoute(
			"auth/checkAuthentication/",
			r.CheckAuth,
			true,
			[]string{},
			httputils.API,
		),
		router.NewPostRoute(
			"auth/login/",
			r.Login,
			false,
			[]string{},
			httputils.API,
		),
		router.NewPostRoute(
			"auth/logout/",
			r.Logout,
			true,
			[]string{},
			httputils.API,
		),
	}
}
