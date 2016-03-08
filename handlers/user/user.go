package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ab22/abcd/httputils"
	"github.com/ab22/abcd/models"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

// Handler defines all methods to be implemented by the User Handler.
type Handler interface {
	FindAllAvailable(context.Context, http.ResponseWriter, *http.Request) error
	FindByID(context.Context, http.ResponseWriter, *http.Request) error
	FindByUsername(context.Context, http.ResponseWriter, *http.Request) error
	Create(context.Context, http.ResponseWriter, *http.Request) error
	Edit(context.Context, http.ResponseWriter, *http.Request) error
	ChangePassword(context.Context, http.ResponseWriter, *http.Request) error
	Delete(context.Context, http.ResponseWriter, *http.Request) error
	GetProfileForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
	ChangePasswordForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
	ChangeEmailForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
	ChangeFullNameForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
	GetPrivilegesForCurrentUser(context.Context, http.ResponseWriter, *http.Request) error
}

// handler structure for the user handlers.
type handler struct {
	services services.Services
}

// NewHandler initializes a new user handler struct.
func NewHandler(s services.Services) *Handler {
	return &Handler{
		services: s,
	}
}

// FindAllAvailable finds all available users.
func (h *Handler) FindAllAvailable(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err         error
		userService = h.services.User()
	)

	type MappedUser struct {
		ID        int       `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Status    int       `json:"status"`
		IsAdmin   bool      `json:"isAdmin"`
		IsTeacher bool      `json:"isTeacher"`
		CreatedAt time.Time `json:"createdAt"`
	}

	users, err := userService.FindAll()
	if err != nil {
		return err
	}

	response := make([]MappedUser, 0, len(users))
	for _, user := range users {
		response = append(response, MappedUser{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Status:    user.Status,
			IsAdmin:   user.IsAdmin,
			IsTeacher: user.IsTeacher,
			CreatedAt: user.CreatedAt,
		})
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}

// FindByID finds a single user by UserID.
func (h *Handler) FindByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err         error
		userID      int
		userService = h.services.User()
		vars        = mux.Vars(r)
	)

	type MappedUser struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Status    int    `json:"status"`
		IsAdmin   bool   `json:"isAdmin"`
		IsTeacher bool   `json:"isTeacher"`
	}

	userID, err = strconv.Atoi(vars["id"])

	if err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return err
	}

	user, err := userService.FindByID(userID)
	if err != nil {
		return err
	} else if user == nil {
		httputils.WriteError(w, http.StatusNotFound, "")
		return nil
	}

	response := &MappedUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}

// FindByUsername finds a User by Username.
func (h *Handler) FindByUsername(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err         error
		userService = h.services.User()
		vars        = mux.Vars(r)
		username    = vars["username"]
	)

	type MappedUser struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Status    int    `json:"status"`
		IsAdmin   bool   `json:"isAdmin"`
		IsTeacher bool   `json:"isTeacher"`
	}

	user, err := userService.FindByUsername(username)
	if err != nil {
		return err
	} else if user == nil {
		httputils.WriteError(w, http.StatusNotFound, "")
		return nil
	}

	response := &MappedUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}

// Edit a user.
func (h *Handler) Edit(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err         error
		userService = h.services.User()

		payload struct {
			ID        int
			Username  string
			Email     string
			FirstName string
			LastName  string
			Status    int
			IsAdmin   bool
			IsTeacher bool
		}
	)

	type Response struct {
		Success      bool   `json:"success"`
		ErrorMessage string `json:"errorMessage"`
	}

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	user := &models.User{
		ID:        payload.ID,
		Username:  payload.Username,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Status:    payload.Status,
		IsAdmin:   payload.IsAdmin,
		IsTeacher: payload.IsTeacher,
	}

	err = userService.Edit(user)
	if err != nil {
		if err == services.ErrDuplicatedUsername {
			return httputils.WriteJSON(w, http.StatusOK, &Response{
				Success:      false,
				ErrorMessage: "El nombre de usuario ya existe!",
			})
		}

		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, &Response{
		Success: true,
	})
}

// Create a user.
func (h *Handler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err         error
		userService = h.services.User()

		payload struct {
			Username  string
			Password  string
			FirstName string
			LastName  string
			Email     string
			IsAdmin   bool
			IsTeacher bool
		}
	)

	type Response struct {
		Success      bool   `json:"success"`
		ErrorMessage string `json:"errorMessage"`
	}

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	user := &models.User{
		Username:  payload.Username,
		Password:  payload.Password,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		IsAdmin:   payload.IsAdmin,
		IsTeacher: payload.IsTeacher,
	}

	err = userService.Create(user)
	if err != nil {
		if err == services.ErrDuplicatedUsername {
			return httputils.WriteJSON(w, http.StatusOK, &Response{
				Success:      false,
				ErrorMessage: "El nombre de usuario ya existe!",
			})
		}

		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, &Response{
		Success: true,
	})
}

// ChangePassword changes a user's password.
func (h *Handler) ChangePassword(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err         error
		userService = h.services.User()

		payload struct {
			UserID      int
			NewPassword string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = userService.ChangePassword(payload.UserID, payload.NewPassword)
	if err != nil && err != services.ErrRecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// Delete user.
func (h *Handler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err         error
		userService = h.services.User()

		payload struct {
			UserID int
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = userService.Delete(payload.UserID)
	if err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// GetProfileForCurrentUser retrieves the logged user's information.
func (h *Handler) GetProfileForCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		userService    = h.services.User()
		sessionData, _ = ctx.Value("sessionData").(*httputils.SessionData)
	)

	type Response struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Status    int    `json:"status"`
		IsAdmin   bool   `json:"isAdmin"`
		IsTeacher bool   `json:"isTeacher"`
	}

	user, err := userService.FindByID(sessionData.UserID)

	if err != nil {
		return err
	} else if user == nil {
		httputils.WriteError(w, http.StatusNotFound, "")
		return nil
	}

	response := &Response{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}

// ChangePasswordForCurrentUser changes the logged user's password.
func (h *Handler) ChangePasswordForCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err            error
		userService    = h.services.User()
		sessionData, _ = ctx.Value("sessionData").(*httputils.SessionData)

		payload struct {
			NewPassword string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = userService.ChangePassword(sessionData.UserID, payload.NewPassword)
	if err != nil && err != services.ErrRecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// ChangeEmailForCurrentUser changes the logged user's email.
func (h *Handler) ChangeEmailForCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err            error
		userService    = h.services.User()
		sessionData, _ = ctx.Value("sessionData").(*httputils.SessionData)

		payload struct {
			NewEmail string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = userService.ChangeEmail(sessionData.UserID, payload.NewEmail)

	if err != nil && err != services.ErrRecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// ChangeFullNameForCurrentUser change the logged user's full name.
func (h *Handler) ChangeFullNameForCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err            error
		userService    = h.services.User()
		sessionData, _ = ctx.Value("sessionData").(*httputils.SessionData)

		payload struct {
			FirstName string
			LastName  string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = userService.ChangeFullName(sessionData.UserID, payload.FirstName, payload.LastName)

	if err != nil && err != services.ErrRecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// GetPrivilegesForCurrentUser returns information about the current user.
// The response includes the IsAdmin and IsTeacher booleans stored in the
// current session.
//
// Changing a user's privilege is something that will not happen very often,
// so in this case, we load them from the session cookie to avoid hitting the
// database everytime. If said user's privileges get changed, then the user
// will have to relog to update the values.
func (h *Handler) GetPrivilegesForCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		sessionData, _ = ctx.Value("sessionData").(*httputils.SessionData)
	)

	type Response struct {
		IsAdmin   bool `json:"isAdmin"`
		IsTeacher bool `json:"isTeacher"`
	}

	response := &Response{
		IsAdmin:   sessionData.IsAdmin,
		IsTeacher: sessionData.IsTeacher,
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}
