package httputils

import (
	"encoding/json"
	"golang.org/x/net/context"
	"io"
	"net/http"
)

type RouteType int

const (
	API RouteType = iota
	Static
)

// APIHandler defines how API handler functions should be defined. ApiHandler
// functions should return any kind of value which will be turned into json
// and an *ApiError.
type APIHandler func(context.Context, http.ResponseWriter, *http.Request) error

func WriteError(w http.ResponseWriter, errMsg string, code int) {
	if errMsg == "" {
		errMsg = http.StatusText(code)
	}

	http.Error(w, errMsg, code)
}

func WriteJSON(w http.ResponseWriter, data interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

func DecodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
