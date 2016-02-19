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

// Handler structure for the user handlers.
type Handler struct {
	services *services.Services
}

// NewHandler initializes a new user handler struct.
func NewHandler(s *services.Services) router.Router {
	return &Handler{
		services: s,
	}
}

// FindAllAvailable finds all available users.
func (h *Handler) FindAllAvailable(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)
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

	users, err := s.User.FindAll()
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
		err    error
		userID int
		s, _   = ctx.Value("services").(*services.Services)
		vars   = mux.Vars(r)
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

	user, err := s.User.FindByID(userID)
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
		err      error
		s, _     = ctx.Value("services").(*services.Services)
		vars     = mux.Vars(r)
		username = vars["username"]
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

	user, err := s.User.FindByUsername(username)
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
		err  error
		s, _ = ctx.Value("services").(*services.Services)

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

	err = s.User.Edit(user)
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
		err  error
		s, _ = ctx.Value("services").(*services.Services)

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

	err = s.User.Create(user)
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
		err  error
		s, _ = ctx.Value("services").(*services.Services)

		payload struct {
			UserID      int
			NewPassword string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = s.User.ChangePassword(payload.UserID, payload.NewPassword)
	if err != nil && err != services.ErrRecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// Delete user.
func (h *Handler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Ruest) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)

		payload struct {
			UserID int
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = s.User.Delete(payload.UserID)
	if err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// GetProfileForCurrentUser retrieves the logged user's information.
func (h *Handler) GetProfileForCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		s, _           = ctx.Value("services").(*services.Services)
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)
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

	user, err := s.User.FindByID(sessionData.UserID)

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
		s, _           = ctx.Value("services").(*services.Services)
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)

		payload struct {
			NewPassword string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = s.User.ChangePassword(sessionData.UserID, payload.NewPassword)
	if err != nil && err != services.ErrRecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// ChangeEmailForCurrentUser changes the logged user's email.
func (h *Handler) ChangeEmailForCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err            error
		s, _           = ctx.Value("services").(*services.Services)
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)

		payload struct {
			NewEmail string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = s.User.ChangeEmail(sessionData.UserID, payload.NewEmail)

	if err != nil && err != services.ErrRecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}

// ChangeFullNameForCurrentUser change the logged user's full name.
func (h *Handler) ChangeFullNameForCurrentUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err            error
		s, _           = ctx.Value("services").(*services.Services)
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)

		payload struct {
			FirstName string
			LastName  string
		}
	)

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	err = s.User.ChangeFullName(sessionData.UserID, payload.FirstName, payload.LastName)

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
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)
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
