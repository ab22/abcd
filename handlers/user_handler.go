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
		RoleName  string    `json:"roleName"`
		Status    int       `json:"status"`
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
			RoleName:  user.Role.Name,
			Status:    user.Status,
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
		RoleId    int    `json:"roleId"`
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
		RoleId:    user.RoleId,
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
		RoleId    int
		Status    int
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
		RoleId:    payload.RoleId,
		Status:    payload.Status,
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
		RoleId    int
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
		RoleId:    payload.RoleId,
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
		RolName   string `json:"roleName"`
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
		RolName:   user.Role.Name,
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
