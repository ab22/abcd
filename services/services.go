package services

import (
	"github.com/jinzhu/gorm"
)

// Services contains all services structures in one whole place.
type Services struct {
	Auth    *authService
	User    *userService
	Student *studentService
}

// NewServices creates a new instance of Services.
func NewServices(db *gorm.DB) *Services {
	s := &Services{}

	s.User = &userService{
		db: db,
	}

	s.Auth = &authService{
		db:          db,
		userService: s.User,
	}

	s.Student = &studentService{
		db: db,
	}

	return s
}
