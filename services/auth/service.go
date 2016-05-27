package auth

import (
	"github.com/ab22/abcd/models"
	"github.com/ab22/abcd/services/user"
	"github.com/jinzhu/gorm"
)

// Contains all of the logic for the systems authentications.
type authService struct {
	db          *gorm.DB
	userService user.Service
}

// NewService creates a new authentication service.
func NewService(db *gorm.DB, userService user.Service) Service {
	return &authService{
		db:          db,
		userService: userService,
	}
}

// Basic username/password authentication. BasicAuth checks if the user exists,
// if the passwords match and if the user's state is set as active.
func (s *authService) BasicAuth(username, password string) (*models.User, error) {
	u, err := s.userService.FindByUsername(username)

	if err != nil {
		return nil, err
	} else if u == nil || u.Status != int(user.Enabled) {
		return nil, nil
	}

	match := s.userService.ComparePasswords([]byte(u.Password), password)

	if !match {
		return nil, nil
	}

	return u, nil
}
