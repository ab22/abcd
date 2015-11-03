package handlers

import (
	"net/http"
	"time"

	_ "github.com/ab22/abcd/models"
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
			CreatedAt: user.CreatedAt,
		})
	}

	return response, nil
}
