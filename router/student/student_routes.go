package student

import (
	"net/http"
	"time"

	"github.com/ab22/abcd/router/httputils"
	"github.com/ab22/abcd/services"
	"golang.org/x/net/context"
)

func (r *studentRouter) FindAllAvailable(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
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

	response := make([]MappedStudent, 0, len(students))

	for _, student := range students {
		response = append(response, MappedStudent{
			Id:        student.Id,
			Email:     student.Email,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Status:    student.Status,
			CreatedAt: student.CreatedAt,
		})
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}
