package user

import "github.com/ab22/abcd/models"

// Service interface describes all functions that must be implemented.
type Service interface {
	FindByID(int) (*models.User, error)
	FindByUsername(string) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	EncryptPassword(string) ([]byte, error)
	ComparePasswords([]byte, string) bool
	FindAll() ([]models.User, error)
	Edit(*models.User) error
	ChangePassword(int, string) error
	Create(*models.User) error
	Delete(int) error
	ChangeEmail(int, string) error
	ChangeFullName(int, string, string) error
}
