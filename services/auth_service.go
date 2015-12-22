package services

import (
	"github.com/ab22/abcd/models"
	"github.com/jinzhu/gorm"
)

// Contains all of the logic for the systems authentications.
type authService struct {
	db          *gorm.DB
	userService *userService
}

// Basic username/password authentication. BasicAuth checks if the user exists,
// if the passwords match and if the user's state is set as active.
func (s *authService) BasicAuth(username, password string) (*models.User, error) {
	user, err := s.userService.FindByUsername(username)

	if err != nil {
		return nil, err
	} else if user == nil || user.Status != int(Enabled) {
		return nil, nil
	}

	match := s.userService.ComparePasswords([]byte(user.Password), password)

	if !match {
		return nil, nil
	}

	return user, nil
}
