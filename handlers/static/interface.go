package static

import (
	"net/http"

	"golang.org/x/net/context"
)

// Handler defines all methods to be implemented by the Static Handler.
type Handler interface {
	Index(context.Context, http.ResponseWriter, *http.Request) error
}
