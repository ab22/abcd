package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ab22/abcd/services"
)

// Contains all handlers in charge of authentication and sessions.
type authHandler struct{}

// We asume that the ValidateAuth decorator called this function
// because the session was validated successfully.
func (h *authHandler) CheckAuth(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	return nil, nil
}

// Basic email/password login.
// Checks:
// 		- User must exist
//		- Passwords match
//		- User's status is Active
//
// If the checks pass, it sets up a session cookie.
func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var (
		err error

		loginForm struct {
			Identifier string
			Password   string
		}
	)

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&loginForm); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	user, err := services.AuthService.BasicAuth(loginForm.Identifier, loginForm.Password)
	if err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	} else if user == nil {
		return nil, &ApiError{
			Error:    nil,
			HttpCode: http.StatusUnauthorized,
			Message:  "Usuario/clave inválidos!",
		}
	}

	session, _ := cookieStore.New(r, sessionCookieName)
	session.Values["data"] = &SessionData{
		UserId:    user.Id,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}
	session.Save(r, w)

	return nil, nil
}

// Basic session logout. If the session cookie is set, it sets its MaxAge=-1
// to delete it, else, well he's already out.
func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	session, err := cookieStore.Get(r, sessionCookieName)

	if err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusUnauthorized,
		}
	}

	session.Options.MaxAge = -1
	session.Save(r, w)
	return nil, nil
}
