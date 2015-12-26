package auth

import (
	"net/http"

	"github.com/ab22/abcd/router/httputils"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

func (r *userRouter) FindAllAvailable(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)
	)

	type MappedStudent struct {
		Id        int       `json:"id"`
		Email     string    `json:"email"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Status    int       `json:"status"`
		CreatedAt time.Time `json:"createdAt"`
	}

	students, err := s.Student.FindAll()
	if err != nil {
		return err
	}

	response := make([]MappedUser, 0, len(students))

	for _, student := range students {
		response = append(response, MappedStudent{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		})
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}
