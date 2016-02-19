package routes

import (
	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/handlers/auth"
	"github.com/ab22/abcd/handlers/static"
	"github.com/ab22/abcd/handlers/user"
	"github.com/ab22/abcd/services"
)

// Routes contains all Template and API routes for the application.
type Routes struct {
	TemplateRoutes []Route
	APIRoutes      []Route
}

// NewRoutes creates a new Router instance and initializes all template
// and API Routes.
func NewRoutes(cfg *config.Config, services *services.Services) *Routes {
	var (
		staticHandler = static.NewHandler(cfg)
		authHandler   = auth.NewHandler(services)
		userHandler   = user.NewHandler(services)

		r = &Routes{
			TemplateRoutes: []Route{
				&route{
					pattern:       "/",
					method:        "GET",
					handlerFunc:   staticHandler.Index,
					requiresAuth:  false,
					requiredRoles: []string{},
				},
			},
			APIRoutes: []Route{
				&route{
					pattern:       "auth/checkAuthentication/",
					method:        "POST",
					handlerFunc:   authHandler.CheckAuth,
					requiresAuth:  true,
					requiredRoles: []string{},
				},
				&route{
					pattern:       "auth/login/",
					method:        "POST",
					handlerFunc:   authHandler.Login,
					requiresAuth:  false,
					requiredRoles: []string{},
				},
				&route{
					pattern:       "auth/logout/",
					method:        "POST",
					handlerFunc:   authHandler.Logout,
					requiresAuth:  true,
					requiredRoles: []string{},
				},
			},
		}
	)

	return r
}
