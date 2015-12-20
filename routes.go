package main

import (
	"net/http"

	"github.com/ab22/abcd/handlers"
)

// Defines an API route with an ApiHandler type as HandlerFunc instead
// of a http.HandleFunc.
type ApiRoute struct {
	Pattern       string
	Method        string
	HandlerFunc   handlers.ApiHandler
	RequiresAuth  bool
	RequiredRoles []string
}

// Simple Template route structure.
type TemplateRoute struct {
	Pattern      string
	Method       string
	HandlerFunc  http.HandlerFunc
	RequiresAuth bool
}

// Routes contains a list of template routes and a list of api routes
// to be registered with middleware as needed.
// Template routes are just routes that serve html templates or
// static files. These might be removed by a nginx http server in the future.
// Api Routes receive and respond json.
type Routes struct {
	TemplateRoutes []TemplateRoute
	ApiRoutes      []ApiRoute
}

var routes = Routes{
	TemplateRoutes: []TemplateRoute{
		{
			Pattern:      "/",
			Method:       "GET",
			HandlerFunc:  handlers.StaticHandler.IndexHandler,
			RequiresAuth: false,
		},
	},
	ApiRoutes: []ApiRoute{
		{
			Pattern:       "auth/checkAuthentication/",
			Method:        "POST",
			HandlerFunc:   handlers.AuthHandler.CheckAuth,
			RequiresAuth:  true,
			RequiredRoles: []string{},
		},
		{
			Pattern:       "auth/login/",
			Method:        "POST",
			HandlerFunc:   handlers.AuthHandler.Login,
			RequiresAuth:  false,
			RequiredRoles: []string{},
		},
		{
			Pattern:       "auth/logout/",
			Method:        "POST",
			HandlerFunc:   handlers.AuthHandler.Logout,
			RequiresAuth:  true,
			RequiredRoles: []string{},
		},
		{
			Pattern:       "user/findAll/",
			Method:        "GET",
			HandlerFunc:   handlers.UserHandler.FindAllAvailable,
			RequiresAuth:  true,
			RequiredRoles: []string{"ADMIN"},
		},
		{
			Pattern:       "user/findById/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.FindById,
			RequiresAuth:  true,
			RequiredRoles: []string{"ADMIN"},
		},
		{
			Pattern:       "user/findByUsername/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.FindByUsername,
			RequiresAuth:  true,
			RequiredRoles: []string{"ADMIN"},
		},
		{
			Pattern:       "user/edit/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.Edit,
			RequiresAuth:  true,
			RequiredRoles: []string{"ADMIN"},
		},
		{
			Pattern:       "user/create/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.Create,
			RequiresAuth:  true,
			RequiredRoles: []string{"ADMIN"},
		},
		{
			Pattern:       "user/changePassword/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.ChangePassword,
			RequiresAuth:  true,
			RequiredRoles: []string{"ADMIN"},
		},
		{
			Pattern:       "user/delete/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.Delete,
			RequiresAuth:  true,
			RequiredRoles: []string{"ADMIN"},
		},
		{
			Pattern:       "user/profile/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.GetProfileForCurrentUser,
			RequiresAuth:  true,
			RequiredRoles: []string{},
		},
		{
			Pattern:       "user/current/changePassword/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.ChangePasswordForCurrentUser,
			RequiresAuth:  true,
			RequiredRoles: []string{},
		},
		{
			Pattern:       "user/current/changeEmail/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.ChangeEmailForCurrentUser,
			RequiresAuth:  true,
			RequiredRoles: []string{},
		},
		{
			Pattern:       "user/current/changeFullName/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.ChangeFullNameForCurrentUser,
			RequiresAuth:  true,
			RequiredRoles: []string{},
		},
		{
			Pattern:       "user/current/privileges/",
			Method:        "POST",
			HandlerFunc:   handlers.UserHandler.GetPrivilegesForCurrentUser,
			RequiresAuth:  true,
			RequiredRoles: []string{},
		},
	},
}
