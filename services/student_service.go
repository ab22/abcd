package services

import (
	"strings"

	"github.com/ab22/abcd/models"
	"github.com/jinzhu/gorm"
)

// StudentService interface describes all functions that must be implemented.
type StudentService interface {
	FindAll() ([]models.Student, error)
	FindByID(int) (*models.Student, error)
	FindByIDNumber(string) (*models.Student, error)
	Create(*models.Student) error
	Edit(*models.Student) error
}

// Contains all logic to handle students in the database.
type studentService struct {
	db *gorm.DB
}

// sanitizeIDNumber trims spaces in the string.
func (s *studentService) sanitizeIDNumber(idNumber string) string {
	return strings.Trim(idNumber, " ")
}

// Finds all active students in the database.
func (s *studentService) FindAll() ([]models.Student, error) {
	var (
		students []models.Student
		err      error
	)

	err = s.db.
		Order("id asc").
		Find(&students).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}

		return students, nil
	}

	return students, nil
}

// Search student by id
func (s *studentService) FindByID(studentID int) (*models.Student, error) {
	student := &models.Student{}

	err := s.db.
		Where("id = ?", studentID).
		First(student).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}

		return nil, nil
	}

	return student, nil
}

// Search student by Id Number/Passport.
func (s *studentService) FindByIDNumber(idNumber string) (*models.Student, error) {
	student := &models.Student{}
	idNumber = s.sanitizeIDNumber(idNumber)

	err := s.db.
		Where("id_number = ?", idNumber).
		First(student).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}

		return nil, nil
	}

	return student, nil
}

// Creates a new student.
func (s *studentService) Create(student *models.Student) error {
	var err error
	student.IDNumber = s.sanitizeIDNumber(student.IDNumber)

	result, err := s.FindByIDNumber(student.IDNumber)

	if err != nil {
		return err
	} else if result != nil {
		return ErrDuplicatedStudentIDNumber
	}

	return s.db.Create(student).Error
}

// Edit an existing student.
func (s *studentService) Edit(newStudent *models.Student) error {
	student, err := s.FindByID(newStudent.ID)

	if err != nil {
		return err
	} else if student == nil {
		return nil
	}

	if student.IDNumber != newStudent.IDNumber {
		duplicateUser, err := s.FindByIDNumber(newStudent.IDNumber)

		if err != nil {
			return err
		} else if duplicateUser != nil {
			return ErrDuplicatedStudentIDNumber
		}
	}

	student.IDNumber = newStudent.IDNumber
	student.FirstName = newStudent.FirstName
	student.LastName = newStudent.LastName
	student.Email = newStudent.Email
	student.Status = newStudent.Status
	student.PlaceOfBirth = newStudent.PlaceOfBirth
	student.Address = newStudent.Address
	student.Birthdate = newStudent.Birthdate
	student.Gender = newStudent.Gender
	student.Nationality = newStudent.Nationality
	student.PhoneNumber = newStudent.PhoneNumber

	return s.db.Save(&student).Error
}
