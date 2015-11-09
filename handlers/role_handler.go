package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ab22/abcd/models"
	"github.com/ab22/abcd/services"
)

// Contains all function handlers for user roles.
type roleHandler struct{}

// Returns all avaialble roles in the database.
func (h *roleHandler) FindAll(w http.ResponseWriter, r *http.Request) (interface{}, *ApiError) {
	var roles []models.Role
	var err error
	type MappedRole struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	roles, err = services.RoleService.FindAll()
	if err != nil {
		return nil, &ApiError{
			Error:    err,
			HttpCode: http.StatusInternalServerError,
		}
	}

	response := make([]MappedRole, 0, len(roles))
	for _, role := range roles {
		response = append(response, MappedUser{
			Id:   role.Id,
			Name: role.Name,
		})
	}

	return response, nil
}
