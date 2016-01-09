package student

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ab22/abcd/models"
	"github.com/ab22/abcd/router/httputils"
	"github.com/ab22/abcd/services"
	"github.com/gorilla/mux"
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

// Find Student by id.
func (r *studentRouter) FindById(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err       error
		studentId int
		s, _      = ctx.Value("services").(*services.Services)
		vars      = mux.Vars(req)
	)

	type MappedStudent struct {
		Id        int       `json:"id"`
		Email     string    `json:"email"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Status    int       `json:"status"`
		CreatedAt time.Time `json:"createdAt"`
	}

	studentId, err = strconv.Atoi(vars["id"])

	if err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return err
	}

	student, err := s.Student.FindById(studentId)
	if err != nil {
		return err
	} else if student == nil {
		httputils.WriteError(w, http.StatusNotFound, "")
		return nil
	}

	response := &MappedStudent{
		Id:        student.Id,
		Email:     student.Email,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Status:    student.Status,
		CreatedAt: student.CreatedAt,
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}

//Create a student.
func (r *studentRouter) Create(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)

		payload struct {
			FirstName string
			LastName  string
			Status    int
		}
	)

	type Response struct {
		Success      bool   `json:"success"`
		ErrorMessage string `json:"errorMessage"`
	}

	if err = httputils.DecodeJSON(req.Body, &payload); err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	student := &models.Student{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Status:    payload.Status,
	}

	err = s.Student.Create(student)
	if err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, nil)
}
