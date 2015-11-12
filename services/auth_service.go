package services

import (
	"github.com/ab22/abcd/models"
)

// Contains all of the logic for the systems authentications.
type authService struct{}

// Basic email/password authentication. BasicAuth checks if the user exists,
// checks if the passwords match and if the user's state is active.
func (s *authService) BasicAuth(username, password string) (*models.User, error) {
	user, err := UserService.FindByUsername(username)
	if err != nil {
		return nil, err
	} else if user == nil || user.Status != int(Enabled) {
		return nil, nil
	}

	match := UserService.ComparePasswords([]byte(user.Password), password)
	if !match {
		return nil, nil
	}

	return user, nil
}
