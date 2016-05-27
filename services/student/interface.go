package student

import "github.com/ab22/abcd/models"

// Service interface describes all functions that must be implemented.
type Service interface {
	FindAll() ([]models.Student, error)
	FindByID(int) (*models.Student, error)
	FindByIDNumber(string) (*models.Student, error)
	Create(*models.Student) error
	Edit(*models.Student) error
}
