package auth

import (
	"net/http"
	"time"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/httputils"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

// Handler defines all methods to be implemented by the Auth Handler.
type Handler interface {
	CheckAuth(context.Context, http.ResponseWriter, *http.Request) error
	Login(context.Context, http.ResponseWriter, *http.Request) error
	Logout(context.Context, http.ResponseWriter, *http.Request) error
}

// Handler structure for the auth handler.
type handler struct {
	services services.Services
}

// NewHandler initializes an auth handler struct.
func NewHandler(s services.Services) Handler {
	return &handler{
		services: s,
	}
}

// CheckAuth asumes that the ValidateAuth decorator called this function
// because the session was validated successfully.
func (h *handler) CheckAuth(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Login does basic email/password login.
// Checks:
// 		- User must exist
//		- Passwords match
//		- User's status is Active
//
// If the checks pass, it sets up a session cookie.
func (h *handler) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		authService = h.services.Auth()
		cookieStore = ctx.Value("cookieStore").(*sessions.CookieStore)
		cfg         = ctx.Value("config").(*config.Config)
		err         error

		loginForm struct {
			Identifier string
			Password   string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &loginForm); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	user, err := authService.BasicAuth(loginForm.Identifier, loginForm.Password)

	if err != nil {
		httputils.WriteError(w, http.StatusInternalServerError, "")
		return nil
	} else if user == nil {
		httputils.WriteError(w, http.StatusUnauthorized, "Usuario/clave inválidos")
		return nil
	}

	session, _ := cookieStore.New(r, cfg.SessionCookieName)
	session.Values["data"] = &httputils.SessionData{
		UserID:    user.ID,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
		ExpiresAt: time.Now().Add(cfg.SessionLifeTime),
	}

	return session.Save(r, w)
}

// Logout does basic session logout.
func (h *handler) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		cookieStore = ctx.Value("cookieStore").(*sessions.CookieStore)
		cfg         = ctx.Value("config").(*config.Config)
		err         error
	)
	session, err := cookieStore.Get(r, cfg.SessionCookieName)

	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}
