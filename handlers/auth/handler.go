package auth

import (
	"net/http"
	"time"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/httputils"
	"github.com/ab22/abcd/services/auth"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

// Handler structure for the auth handler.
type handler struct {
	cfg         *config.Config
	authService auth.Service
}

// NewHandler initializes an auth handler struct.
func NewHandler(cfg *config.Config, authService auth.Service) Handler {
	return &handler{
		cfg:         cfg,
		authService: authService,
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
		cookieStore = ctx.Value("cookieStore").(*sessions.CookieStore)
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

	user, err := h.authService.BasicAuth(loginForm.Identifier, loginForm.Password)

	if err != nil {
		httputils.WriteError(w, http.StatusInternalServerError, "")
		return nil
	} else if user == nil {
		httputils.WriteError(w, http.StatusUnauthorized, "Usuario/clave inválidos")
		return nil
	}

	session, _ := cookieStore.New(r, h.cfg.SessionCookieName)
	session.Values["data"] = &httputils.SessionData{
		UserID:    user.ID,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
		ExpiresAt: time.Now().Add(h.cfg.SessionLifeTime),
	}

	return session.Save(r, w)
}

// Logout does basic session logout.
func (h *handler) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		cookieStore = ctx.Value("cookieStore").(*sessions.CookieStore)
		err         error
	)
	session, err := cookieStore.Get(r, h.cfg.SessionCookieName)

	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}
