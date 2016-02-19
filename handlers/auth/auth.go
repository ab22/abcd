package auth

import (
	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/httputils"
	"golang.org/x/net/context"
)

// Handler structure for the auth handler.
type Handler struct {
	services *services.Services
}

// NewHandler initializes an auth handler struct.
func NewHandler(s *services.Services) *Handler {
	return &Handler{
		services: s,
	}
}

// CheckAuth asumes that the ValidateAuth decorator called this function
// because the session was validated successfully.
func (h *Handler) CheckAuth(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Login does basic email/password login.
// Checks:
// 		- User must exist
//		- Passwords match
//		- User's status is Active
//
// If the checks pass, it sets up a session cookie.
func (h *Handler) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		authService = h.services.Auth
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

	user, err := s.Auth.BasicAuth(loginForm.Identifier, loginForm.Password)
	if err != nil {
		httputils.WriteError(w, http.StatusInternalServerError, "")
		return nil
	} else if user == nil {
		httputils.WriteError(w, http.StatusUnauthorized, "Usuario/clave inválidos")
		return nil
	}

	session, _ := cookieStore.New(r, router.SessionCookieName)
	session.Values["data"] = &router.SessionData{
		UserID:    user.ID,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
		ExpiresAt: time.Now().Add(cfg.SessionLifeTime),
	}

	return session.Save(r, w)
}

// Logout does basic session logout.
func (h *Handler) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		cookieStore = ctx.Value("cookieStore").(*sessions.CookieStore)
		err         error
	)
	session, err := cookieStore.Get(r, router.SessionCookieName)

	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}
