package handlers

import (
	"encoding/gob"
	"html/template"
	"net/http"

	"github.com/ab22/abcd/config"
	"github.com/gorilla/sessions"
)

const (
	sessionCookieName = "_session" // Name used for the session cookie.
)

// Global Variables
var (
	cachedTemplates *template.Template    // Cached templates are stored globally.
	cookieStore     *sessions.CookieStore // Global cookie store/provider.

	// Define and export global handlers to encapsulate handler's
	// functions and avoid name collisions.
	StaticHandler = staticHandler{}
)

// Describes the data that every session cookie must store.
type SessionData struct {
	UserId int
	Email  string
}

// Returned by the ApiHandler functions. Contains information about the error
// encountered, the http code to respond as and a string message.
type ApiError struct {
	Error    error
	HttpCode int
	Message  string
}

// ApiHandler defines how API handler functions should be defined. ApiHandler
// functions should return any kind of value which will be turned into json
// and an *ApiError.
type ApiHandler func(http.ResponseWriter, *http.Request) (interface{}, *ApiError)

// Initializes all global variables such as the instance for session storage
// and loads cached templates.
func Initialize() {
	secretKey := config.EnvVariables.App.Secret
	adminAppPath := config.EnvVariables.App.Frontend.Admin

	gob.Register(&SessionData{})

	cookieStore = sessions.NewCookieStore([]byte(secretKey))
	cachedTemplates = template.Must(template.ParseFiles(adminAppPath + "/index.html"))
}
