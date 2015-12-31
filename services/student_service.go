package services

import (
	"github.com/ab22/abcd/models"
	"github.com/jinzhu/gorm"
)

// Contains all logic to handle students in the database.
type studentService struct {
	db *gorm.DB
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
func (s *studentService) FindById(studnetId int) (*models.Student, error) {
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
