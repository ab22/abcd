package student

import (
	"net/http"

	"golang.org/x/net/context"
)

// Handler defines all methods to be implemented by the Student Handler.
type Handler interface {
	FindAllAvailable(context.Context, http.ResponseWriter, *http.Request) error
	FindByID(context.Context, http.ResponseWriter, *http.Request) error
	Create(context.Context, http.ResponseWriter, *http.Request) error
	Edit(context.Context, http.ResponseWriter, *http.Request) error
	FindByIDNumber(context.Context, http.ResponseWriter, *http.Request) error
}
