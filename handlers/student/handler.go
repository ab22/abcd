package student

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/httputils"
	"github.com/ab22/abcd/models"
	"github.com/ab22/abcd/services"
	"github.com/ab22/abcd/services/student"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

// handler structure for the student handlers.
type handler struct {
	cfg            *config.Config
	studentService student.Service
}

// NewHandler initializes a new student handler struct.
func NewHandler(cfg *config.Config, studentService student.Service) Handler {
	return &handler{
		cfg:            cfg,
		studentService: studentService,
	}
}

// FindAllAvailable returns all available students.
func (h *handler) FindAllAvailable(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var err error

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

	students, err := h.studentService.FindAll()
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

// FindByID finds a Student by Id.
func (h *handler) FindByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err       error
		studentID int
		vars      = mux.Vars(r)
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

	student, err := h.studentService.FindByID(studentID)
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

// Create a student.
func (h *handler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err error

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

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
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

	err = h.studentService.Create(student)
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
func (h *handler) Edit(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err error

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

	if err = httputils.DecodeJSON(r.Body, &payload); err != nil {
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

	err = h.studentService.Edit(student)
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

// FindByIDNumber finds student by honduran Id number or passport number.
func (h *handler) FindByIDNumber(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		err      error
		idNumber string
		vars     = mux.Vars(r)
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

	student, err := h.studentService.FindByIDNumber(idNumber)

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
