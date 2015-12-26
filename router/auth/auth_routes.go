package auth

import (
	"net/http"

	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

// We asume that the ValidateAuth decorator called this function
// because the session was validated successfully.
func (r *authRouter) CheckAuth(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	return nil
}

// Basic email/password login.
// Checks:
// 		- User must exist
//		- Passwords match
//		- User's status is Active
//
// If the checks pass, it sets up a session cookie.
func (r *authRouter) Login(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		s, _        = ctx.Value("services").(*services.Services)
		cookieStore = ctx.Value("cookieStore").(*sessions.CookieStore)
		err         error

		loginForm struct {
			Identifier string
			Password   string
		}
	)

	if err = httputils.DecodeJSON(req.Body, &loginForm); err != nil {
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

	session, _ := cookieStore.New(req, router.SessionCookieName)
	session.Values["data"] = &router.SessionData{
		UserId:    user.Id,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}
	session.Save(req, w)

	return nil
}

// Basic session logout. If the session cookie is set, it sets its MaxAge=-1
// to delete it, else, well he's already out.
func (r *authRouter) Logout(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		cookieStore = ctx.Value("cookieStore").(*sessions.CookieStore)
		err         error
	)
	session, err := cookieStore.Get(req, router.SessionCookieName)

	if err != nil {
		httputils.WriteError(w, http.StatusUnauthorized, "")
		return err
	}

	session.Options.MaxAge = -1
	session.Save(req, w)
	return nil
}
