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

// Returns all available students.
func (r *studentRouter) FindAllAvailable(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)
	)

	type MappedStudent struct {
		ID        int       `json:"id"`
		IDNumber  string    `json:"idNumber"`
		Email     string    `json:"email"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Status    int       `json:"status"`
		Birthdate time.Time `json:"birthdate"`
		CreatedAt time.Time `json:"createdAt"`
	}

	students, err := s.Student.FindAll()
	if err != nil {
		return err
	}

	response := make([]MappedStudent, 0, len(students))

	for _, student := range students {
		response = append(response, MappedStudent{
			ID:        student.ID,
			IDNumber:  student.IDNumber,
			Email:     student.Email,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Status:    student.Status,
			Birthdate: student.Birthdate,
			CreatedAt: student.CreatedAt,
		})
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}

// Find Student by Id.
func (r *studentRouter) FindByID(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err       error
		studentID int
		s, _      = ctx.Value("services").(*services.Services)
		vars      = mux.Vars(req)
	)

	type MappedStudent struct {
		ID        int       `json:"id"`
		Email     string    `json:"email"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Status    int       `json:"status"`
		CreatedAt time.Time `json:"createdAt"`
	}

	studentID, err = strconv.Atoi(vars["id"])

	if err != nil {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return nil
	}

	student, err := s.Student.FindByID(studentID)
	if err != nil {
		return err
	} else if student == nil {
		httputils.WriteError(w, http.StatusNotFound, "")
		return nil
	}

	response := &MappedStudent{
		ID:        student.ID,
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
			IDNumber     string
			FirstName    string
			LastName     string
			Email        string
			PlaceOfBirth string
			Address      string
			Birthdate    time.Time
			Gender       bool
			Nationality  string
			PhoneNumber  string
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
		IDNumber:     payload.IDNumber,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		PlaceOfBirth: payload.PlaceOfBirth,
		Address:      payload.Address,
		Birthdate:    payload.Birthdate,
		Gender:       payload.Gender,
		Nationality:  payload.Nationality,
		PhoneNumber:  payload.PhoneNumber,
	}

	err = s.Student.Create(student)
	if err != nil {
		if err == services.ErrDuplicatedStudentIDNumber {
			return httputils.WriteJSON(w, http.StatusOK, &Response{
				Success:      false,
				ErrorMessage: "El número de cédula o pasaporte ya existe!",
			})
		}

		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, &Response{
		Success: true,
	})
}

// Edit a student.
func (r *studentRouter) Edit(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err  error
		s, _ = ctx.Value("services").(*services.Services)

		payload struct {
			ID           int
			IDNumber     string
			FirstName    string
			LastName     string
			Email        string
			PlaceOfBirth string
			Address      string
			Birthdate    time.Time
			Gender       bool
			Nationality  string
			PhoneNumber  string
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
		ID:           payload.ID,
		IDNumber:     payload.IDNumber,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		PlaceOfBirth: payload.PlaceOfBirth,
		Address:      payload.Address,
		Birthdate:    payload.Birthdate,
		Gender:       payload.Gender,
		Nationality:  payload.Nationality,
		PhoneNumber:  payload.PhoneNumber,
	}

	err = s.Student.Edit(student)
	if err != nil {
		if err == services.ErrDuplicatedStudentIDNumber {
			return httputils.WriteJSON(w, http.StatusOK, &Response{
				Success:      false,
				ErrorMessage: "El número de cédula o pasaporte ya existe!",
			})
		}

		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, &Response{
		Success: true,
	})
}

// Find student by honduran Id number or passport number.
func (r *studentRouter) FindByIDNumber(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	var (
		err      error
		idNumber string
		s, _     = ctx.Value("services").(*services.Services)
		vars     = mux.Vars(req)
	)

	type MappedStudent struct {
		ID           int       `json:"id"`
		IDNumber     string    `json:"idNumber"`
		FirstName    string    `json:"firstName"`
		LastName     string    `json:"lastName"`
		Email        string    `json:"email"`
		Status       int       `json:"status"`
		PlaceOfBirth string    `json:"placeOfBirth"`
		Address      string    `json:"address"`
		Birthdate    time.Time `json:"birthdate"`
		Gender       bool      `json:"gender"`
		Nationality  string    `json:"nationality"`
		PhoneNumber  string    `json:"phoneNumber"`
	}

	idNumber = vars["idNumber"]

	if idNumber == "" {
		httputils.WriteError(w, http.StatusBadRequest, "")
		return err
	}

	student, err := s.Student.FindByIDNumber(idNumber)

	if err != nil {
		return err
	} else if student == nil {
		httputils.WriteError(w, http.StatusNotFound, "")
		return nil
	}

	response := &MappedStudent{
		ID:           student.ID,
		IDNumber:     student.IDNumber,
		FirstName:    student.FirstName,
		LastName:     student.LastName,
		Email:        student.Email,
		Status:       student.Status,
		PlaceOfBirth: student.PlaceOfBirth,
		Address:      student.Address,
		Birthdate:    student.Birthdate,
		Gender:       student.Gender,
		Nationality:  student.Nationality,
		PhoneNumber:  student.PhoneNumber,
	}

	return httputils.WriteJSON(w, http.StatusOK, response)
}
