package auth

import (
	"net/http"

	"golang.org/x/net/context"
)

// Handler defines all methods to be implemented by the Auth Handler.
type Handler interface {
	CheckAuth(context.Context, http.ResponseWriter, *http.Request) error
	Login(context.Context, http.ResponseWriter, *http.Request) error
	Logout(context.Context, http.ResponseWriter, *http.Request) error
}
