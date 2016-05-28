package user

import (
	"net/http"

	"golang.org/x/net/context"
)

// Handler defines all methods to be implemented by the User Handler.
type Handler interface {
	FindAllAvailable(context.Context, http.ResponseWriter, *http.Request) error
	FindByID(context.Context, http.ResponseWriter, *http.Request) error
	FindByUsername(context.Context, http.ResponseWriter, *http.Request) error
	Create(context.Context, http.ResponseWriter, *http.Request) error
	Edit(context.Context, http.ResponseWriter, *http.Request) error
	ChangePassword(context.Context, http.ResponseWriter, *http.Request) error
	Delete(context.Context, http.ResponseWriter, *http.Request) error
	GetProfileForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
	ChangePasswordForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
	ChangeEmailForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
	ChangeFullNameForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
	GetPrivilegesForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
}
