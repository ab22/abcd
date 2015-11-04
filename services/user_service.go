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

// Search user by id
func (s *userService) FindById(userId int) (*models.User, error) {
	user := &models.User{}

	err := db.
		Where("id = ?", userId).
		First(user).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return nil, nil
	}

	return user, nil
}

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

	err = db.Order("username asc").Find(&users).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return nil, err
		}

		return users, nil
	}

	return users, nil
}

// Edit user modifies the basic user's models attributes. The function checks
// if the username changed and if it needs to check if the username is already
// taken.
func (s *userService) Edit(newUser *models.User) error {
	user, err := s.FindById(newUser.Id)

	if err != nil {
		return err
	} else if user == nil {
		return nil
	}

	if user.Username != newUser.Username {
		duplicateUser, err := s.FindByUsername(newUser.Username)
		if err != nil {
			return err
		} else if duplicateUser != nil {
			return DuplicateUsernameError
		}
	}

	user.Username = newUser.Username
	user.Email = newUser.Email
	user.FirstName = newUser.FirstName
	user.LastName = newUser.LastName
	user.Status = newUser.Status

	return db.Save(&user).Error
}

// ChangePassword finds a user in the database by userId and changes it's
// password.
func (s *userService) ChangePassword(userId int, password string) error {
	encryptedPassword, err := s.EncryptPassword(password)
	if err != nil {
		return err
	}

	err = db.
		Table("users").
		Where("id = ?", userId).
		Update("password", string(encryptedPassword)).Error

	if err != nil {
		if err != gorm.RecordNotFound {
			return err
		}
	} else if db.RowsAffected == 0 {
		return RecordNotFound
	}

	return nil
}
