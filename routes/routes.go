package routes

import (
	"github.com/jinzhu/gorm"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/handlers/auth"
	"github.com/ab22/abcd/handlers/static"
	"github.com/ab22/abcd/handlers/student"
	"github.com/ab22/abcd/handlers/user"
	authservices "github.com/ab22/abcd/services/auth"
	studentservices "github.com/ab22/abcd/services/student"
	userservices "github.com/ab22/abcd/services/user"
)

// Routes contains all Template and API routes for the application.
type Routes struct {
	TemplateRoutes []Route
	APIRoutes      []Route
}

// NewRoutes creates a new Router instance and initializes all template
// and API Routes.
func NewRoutes(cfg *config.Config, db *gorm.DB) *Routes {
	var (
		userService    = userservices.NewService(db)
		authService    = authservices.NewService(db, userService)
		studentService = studentservices.NewService(db)

		staticHandler  = static.NewHandler(cfg)
		authHandler    = auth.NewHandler(cfg, authService)
		userHandler    = user.NewHandler(cfg, userService)
		studentHandler = student.NewHandler(cfg, studentService)

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
					requiresAuth:  false,
					requiredRoles: []string{},
				},
				&route{
					pattern:       "user/findAll/",
					method:        "GET",
					handlerFunc:   userHandler.FindAllAvailable,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "user/findById/{id:[0-9]+}",
					method:        "GET",
					handlerFunc:   userHandler.FindByID,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "user/findByUsername/{username:[0-9a-zA-Z]+}",
					method:        "GET",
					handlerFunc:   userHandler.FindByUsername,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "user/edit/",
					method:        "POST",
					handlerFunc:   userHandler.Edit,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "user/create/",
					method:        "POST",
					handlerFunc:   userHandler.Create,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "user/changePassword/",
					method:        "POST",
					handlerFunc:   userHandler.ChangePassword,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "user/delete/",
					method:        "POST",
					handlerFunc:   userHandler.Delete,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "user/profile/",
					method:        "GET",
					handlerFunc:   userHandler.GetProfileForCurrentUser,
					requiresAuth:  true,
					requiredRoles: []string{},
				},
				&route{
					pattern:       "user/current/changePassword/",
					method:        "POST",
					handlerFunc:   userHandler.ChangePasswordForCurrentUser,
					requiresAuth:  true,
					requiredRoles: []string{},
				},
				&route{
					pattern:       "user/current/changeEmail/",
					method:        "POST",
					handlerFunc:   userHandler.ChangeEmailForCurrentUser,
					requiresAuth:  true,
					requiredRoles: []string{},
				},
				&route{
					pattern:       "user/current/changeFullName/",
					method:        "POST",
					handlerFunc:   userHandler.ChangeFullNameForCurrentUser,
					requiresAuth:  true,
					requiredRoles: []string{},
				},
				&route{
					pattern:       "user/current/privileges/",
					method:        "GET",
					handlerFunc:   userHandler.GetPrivilegesForCurrentUser,
					requiresAuth:  true,
					requiredRoles: []string{},
				},
				&route{
					pattern:       "student/findAll/",
					method:        "GET",
					handlerFunc:   studentHandler.FindAllAvailable,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "student/findById/{id:[0-9]+}",
					method:        "GET",
					handlerFunc:   studentHandler.FindByID,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "student/edit/",
					method:        "POST",
					handlerFunc:   studentHandler.Edit,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "student/create/",
					method:        "POST",
					handlerFunc:   studentHandler.Create,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
				&route{
					pattern:       "student/findByIdNumber/{idNumber:[0-9A-Za-z-]+}",
					method:        "GET",
					handlerFunc:   studentHandler.FindByIDNumber,
					requiresAuth:  true,
					requiredRoles: []string{"ADMIN"},
				},
			},
		}
	)

	return r
}
