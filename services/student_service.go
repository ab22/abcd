package services

import (
	"strings"

	"github.com/ab22/abcd/models"
	"github.com/jinzhu/gorm"
)

// Contains all logic to handle students in the database.
type studentService struct {
	db *gorm.DB
}

func (s *studentService) SanitizeIdNumber(idNumber string) string {
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
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return students, nil
	}

	return students, nil
}

// Search student by id
func (s *studentService) FindById(studentId int) (*models.Student, error) {
	student := &models.Student{}

	err := s.db.
		Where("id = ?", studentId).
		First(student).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return nil, nil
	}

	return student, nil
}

// Search student by Id Number/Passport.
func (s *studentService) FindByIdNumber(idNumber string) (*models.Student, error) {
	student := &models.Student{}
	idNumber = s.SanitizeIdNumber(idNumber)

	err := s.db.
		Where("id_number = ?", idNumber).
		First(student).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return nil, nil
	}

	return student, nil
}

// Creates a new student.
func (s *studentService) Create(student *models.Student) error {
	var err error
	student.IdNumber = s.SanitizeIdNumber(student.IdNumber)

	result, err := s.FindByIdNumber(student.IdNumber)

	if err != nil {
		return err
	} else if result != nil {
		return DuplicatedStudentIdNumberError
	}

	return s.db.Create(student).Error
}

// Edit an existing student.
func (s *studentService) Edit(newStudent *models.Student) error {
	student, err := s.FindById(newStudent.Id)

	if err != nil {
		return err
	} else if student == nil {
		return nil
	}

	if student.IdNumber != newStudent.IdNumber {
		duplicateUser, err := s.FindByIdNumber(newStudent.IdNumber)

		if err != nil {
			return err
		} else if duplicateUser != nil {
			return DuplicatedStudentIdNumberError
		}
	}

	student.IdNumber = newStudent.IdNumber
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
