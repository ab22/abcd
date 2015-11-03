package services

import (
	"github.com/ab22/abcd/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Contains all of the logic for the User model.
type userService struct{}

// Int type to define statuses for the User model.
type UserStatus int

// Defines all user statuses.
const (
	Active UserStatus = iota
	Disabled
)

// Search user by Username.
func (s *userService) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}

	err := db.
		Where("username = ?", username).
		First(user).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return nil, nil
	}

	return user, nil
}

// Searches for a User by Email.
// Returns *models.User instance if it finds it, or nil otherwise.
func (s *userService) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := db.
		Where("email = ?", email).
		First(user).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return nil, nil
	}

	return user, nil
}

// Encrypts a password with the default password hasher (bcrypt).
// Returns the hashed password []byte and an error.
func (s *userService) EncryptPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

// Compares if the hashed password equals the plain text password.
func (s *userService) ComparePasswords(hashedPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}

// Finds all active users in the database.
func (s *userService) FindAll() ([]models.User, error) {
	var users []models.User
	var err error

	err = db.Find(&users).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return users, nil
	}

	return users, nil
}
