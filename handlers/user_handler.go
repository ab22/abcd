package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ab22/abcd/models"
	"github.com/ab22/abcd/services"
)

// Contains all handlers in charge of requests that require handling of the
// User model.
type userHandler struct{}

func (h *userHandler) FindAllAvailable(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	type MappedUser struct {
		Id        int       `json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Status    int       `json:"status"`
		IsAdmin   bool      `json:"isAdmin"`
		IsTeacher bool      `json:"isTeacher"`
		CreatedAt time.Time `json:"createdAt"`
	}

	users, err := services.UserService.FindAll()
	if err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	response := make([]MappedUser, 0, len(users))
	for _, user := range users {
		response = append(response, MappedUser{
			Id:        user.Id,
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

	return response, nil
}

func (h *userHandler) FindById(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	var payload struct {
		UserId int
	}
	type MappedUser struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Status    int    `json:"status"`
		IsAdmin   bool   `json:"isAdmin"`
		IsTeacher bool   `json:"isTeacher"`
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	user, err := services.UserService.FindById(payload.UserId)
	if err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	} else if user == nil {
		return nil, &ApiError{
			Error:    nil,
			HttpCode: http.StatusNotFound,
		}
	}

	response := &MappedUser{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}

	return response, nil
}

func (h *userHandler) FindByUsername(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	var payload struct {
		Username string
	}
	type MappedUser struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Status    int    `json:"status"`
		IsAdmin   bool   `json:"isAdmin"`
		IsTeacher bool   `json:"isTeacher"`
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	user, err := services.UserService.FindByUsername(payload.Username)
	if err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	} else if user == nil {
		return nil, &ApiError{
			Error:    nil,
			HttpCode: http.StatusNotFound,
		}
	}

	response := &MappedUser{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}

	return response, nil
}

// Edit a user.
func (h *userHandler) Edit(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	var payload struct {
		Id        int
		Username  string
		Email     string
		FirstName string
		LastName  string
		Status    int
		IsAdmin   bool
		IsTeacher bool
	}
	type Response struct {
		Success      bool   `json:"success"`
		ErrorMessage string `json:"errorMessage"`
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	user := &models.User{
		Id:        payload.Id,
		Username:  payload.Username,
		Email:     payload.Email,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Status:    payload.Status,
		IsAdmin:   payload.IsAdmin,
		IsTeacher: payload.IsTeacher,
	}

	err = services.UserService.Edit(user)
	if err != nil {
		if err == services.DuplicateUsernameError {
			return &Response{
				Success:      false,
				ErrorMessage: "El nombre de usuario ya existe!",
			}, nil
		}

		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	return &Response{
		Success: true,
	}, nil
}

//Create a user.
func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	var payload struct {
		Username  string
		Password  string
		FirstName string
		LastName  string
		Email     string
		IsAdmin   bool
		IsTeacher bool
	}
	type Response struct {
		Success      bool   `json:"success"`
		ErrorMessage string `json:"errorMessage"`
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
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

	err = services.UserService.Create(user)
	if err != nil {
		if err == services.DuplicateUsernameError {
			return &Response{
				Success:      false,
				ErrorMessage: "El nombre de usuario ya existe!",
			}, nil
		}

		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	return &Response{
		Success: true,
	}, nil
}

// Change a user's password.
func (h *userHandler) ChangePassword(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	var payload struct {
		UserId      int
		NewPassword string
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	err = services.UserService.ChangePassword(payload.UserId, payload.NewPassword)
	if err != nil && err != services.RecordNotFound {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	return nil, nil
}

// Delete user.
func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	var payload struct {
		UserId int
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	err = services.UserService.Delete(payload.UserId)
	if err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	return nil, nil
}

// Retrieve logged user's information.
func (h *userHandler) GetProfileForCurrentUser(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	session, _ := cookieStore.Get(r, sessionCookieName)
	sessionData := session.Values["data"].(*SessionData)
	type Response struct {
		Id        int    `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Status    int    `json:"status"`
		IsAdmin   bool   `json:"isAdmin"`
		IsTeacher bool   `json:"isTeacher"`
	}

	user, err := services.UserService.FindById(sessionData.UserId)
	if err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	} else if user == nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusNotFound,
		}
	}

	return &Response{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}, nil
}

// Change the logged user's password.
func (h *userHandler) ChangePasswordForCurrentUser(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	session, _ := cookieStore.Get(r, sessionCookieName)
	sessionData := session.Values["data"].(*SessionData)
	var payload struct {
		NewPassword string
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	err = services.UserService.ChangePassword(sessionData.UserId, payload.NewPassword)
	if err != nil && err != services.RecordNotFound {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	return nil, nil
}

// Change the logged user's email.
func (h *userHandler) ChangeEmailForCurrentUser(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	session, _ := cookieStore.Get(r, sessionCookieName)
	sessionData := session.Values["data"].(*SessionData)
	var payload struct {
		NewEmail string
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	err = services.UserService.ChangeEmail(sessionData.UserId, payload.NewEmail)
	if err != nil && err != services.RecordNotFound {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	return nil, nil
}

// Change the logged user's full name.
func (h *userHandler) ChangeFullNameForCurrentUser(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var err error
	session, _ := cookieStore.Get(r, sessionCookieName)
	sessionData := session.Values["data"].(*SessionData)
	var payload struct {
		FirstName string
		LastName  string
	}

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&payload); err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusBadRequest,
		}
	}

	err = services.UserService.ChangeFullName(sessionData.UserId, payload.FirstName, payload.LastName)
	if err != nil && err != services.RecordNotFound {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	return nil, nil
}

// GetUserPrivileges returns information about the current user. The response
// includes the IsAdmin and IsTeacher booleans stored in the current session.
//
// Changing a user's privilege is something that will not happen very often,
// so in this case, we load them from the session cookie to avoid hitting the
// database everytime. If said user's privileges get changed, then the user
// will have to relog to update the values.
func (h *authHandler) GetPrivilegesForCurrentUser(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	session, _ := cookieStore.Get(r, sessionCookieName)
	sessionData := session.Values["data"].(*SessionData)

	type Response struct {
		IsAdmin   bool `json:"isAdmin"`
		IsTeacher bool `json:"isTeacher"`
	}

	return &Response{
		IsAdmin:   sessionData.IsAdmin,
		IsTeacher: sessionData.IsTeacher,
	}, nil
}
