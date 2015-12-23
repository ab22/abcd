package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ab22/abcd/models"
	"github.com/ab22/abcd/router"
	"github.com/ab22/abcd/router/httputils"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func (r *userRouter) FindAllAvailable(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)
	)

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

	users, err := s.User.FindAll()
	if err != nil {
		return err
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

	return httputils.WriteJSON(w, response, http.StatusOK)
}

func (r *userRouter) FindById(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err    error
		userId int
		s, _   = ctx.Value("services").(*services.Services)
		vars   = mux.Vars(req)
	)

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

	userId, err = strconv.Atoi(vars["id"])

	if err != nil {
		httputils.WriteError(w, "", http.StatusBadRequest)
		return nil
	}

	user, err := s.User.FindById(userId)
	if err != nil {
		return err
	} else if user == nil {
		httputils.WriteError(w, "", http.StatusNotFound)
		return nil
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

	return httputils.WriteJSON(w, response, http.StatusOK)
}

func (r *userRouter) FindByUsername(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err      error
		s, _     = ctx.Value("services").(*services.Services)
		vars     = mux.Vars(req)
		username = vars["username"]
	)

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

	user, err := s.User.FindByUsername(username)
	if err != nil {
		return err
	} else if user == nil {
		httputils.WriteError(w, "", http.StatusNotFound)
		return nil
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

	return httputils.WriteJSON(w, response, http.StatusOK)
}

// Edit a user.
func (r *userRouter) Edit(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)

		payload struct {
			Id        int
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

	if err = httputils.DecodeJSON(req.Body, &payload); err != nil {
		httputils.WriteError(w, "", http.StatusBadRequest)
		return nil
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

	err = s.User.Edit(user)
	if err != nil {
		if err == services.DuplicateUsernameError {
			return httputils.WriteJSON(w, &Response{
				Success:      false,
				ErrorMessage: "El nombre de usuario ya existe!",
			}, http.StatusOK)
		}

		return err
	}

	return httputils.WriteJSON(w, &Response{
		Success: true,
	}, http.StatusOK)
}

//Create a user.
func (r *userRouter) Create(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
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

	if err = httputils.DecodeJSON(req.Body, &payload); err != nil {
		httputils.WriteError(w, "", http.StatusBadRequest)
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
		if err == services.DuplicateUsernameError {
			return httputils.WriteJSON(w, &Response{
				Success:      false,
				ErrorMessage: "El nombre de usuario ya existe!",
			}, http.StatusOK)
		}

		return err
	}

	return httputils.WriteJSON(w, &Response{
		Success: true,
	}, http.StatusOK)
}

// Change a user's password.
func (r *userRouter) ChangePassword(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)

		payload struct {
			UserId      int
			NewPassword string
		}
	)

	if err = httputils.DecodeJSON(req.Body, &payload); err != nil {
		httputils.WriteError(w, "", http.StatusBadRequest)
		return nil
	}

	err = s.User.ChangePassword(payload.UserId, payload.NewPassword)
	if err != nil && err != services.RecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, nil, http.StatusOK)
}

// Delete user.
func (r *userRouter) Delete(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)

		payload struct {
			UserId int
		}
	)

	if err = httputils.DecodeJSON(req.Body, &payload); err != nil {
		httputils.WriteError(w, "", http.StatusBadRequest)
		return nil
	}

	err = s.User.Delete(payload.UserId)
	if err != nil {
		return err
	}

	return httputils.WriteJSON(w, nil, http.StatusOK)
}

// Retrieve logged user's information.
func (r *userRouter) GetProfileForCurrentUser(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		s, _           = ctx.Value("services").(*services.Services)
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)
	)

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

	user, err := s.User.FindById(sessionData.UserId)

	if err != nil {
		return err
	} else if user == nil {
		httputils.WriteError(w, "", http.StatusNotFound)
		return nil
	}

	response := &Response{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		IsAdmin:   user.IsAdmin,
		IsTeacher: user.IsTeacher,
	}

	return httputils.WriteJSON(w, response, http.StatusOK)
}

// Change the logged user's password.
func (r *userRouter) ChangePasswordForCurrentUser(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err            error
		s, _           = ctx.Value("services").(*services.Services)
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)

		payload struct {
			NewPassword string
		}
	)

	if err = httputils.DecodeJSON(req.Body, &payload); err != nil {
		httputils.WriteError(w, "", http.StatusBadRequest)
		return nil
	}

	err = s.User.ChangePassword(sessionData.UserId, payload.NewPassword)
	if err != nil && err != services.RecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, nil, http.StatusOK)
}

// Change the logged user's email.
func (r *userRouter) ChangeEmailForCurrentUser(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err            error
		s, _           = ctx.Value("services").(*services.Services)
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)

		payload struct {
			NewEmail string
		}
	)

	if err = httputils.DecodeJSON(req.Body, &payload); err != nil {
		httputils.WriteError(w, "", http.StatusBadRequest)
		return nil
	}

	err = s.User.ChangeEmail(sessionData.UserId, payload.NewEmail)

	if err != nil && err != services.RecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, nil, http.StatusOK)
}

// Change the logged user's full name.
func (r *userRouter) ChangeFullNameForCurrentUser(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err            error
		s, _           = ctx.Value("services").(*services.Services)
		sessionData, _ = ctx.Value("sessionData").(*router.SessionData)

		payload struct {
			FirstName string
			LastName  string
		}
	)

	if err = httputils.DecodeJSON(req.Body, &payload); err != nil {
		httputils.WriteError(w, "", http.StatusBadRequest)
		return nil
	}

	err = s.User.ChangeFullName(sessionData.UserId, payload.FirstName, payload.LastName)

	if err != nil && err != services.RecordNotFound {
		return err
	}

	return httputils.WriteJSON(w, nil, http.StatusOK)
}

// GetUserPrivileges returns information about the current user. The response
// includes the IsAdmin and IsTeacher booleans stored in the current session.
//
// Changing a user's privilege is something that will not happen very often,
// so in this case, we load them from the session cookie to avoid hitting the
// database everytime. If said user's privileges get changed, then the user
// will have to relog to update the values.
func (r *userRouter) GetPrivilegesForCurrentUser(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
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

	return httputils.WriteJSON(w, response, http.StatusOK)
}
