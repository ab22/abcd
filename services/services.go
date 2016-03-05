package services

import (
	"github.com/jinzhu/gorm"
)

// Services interface defines all methods that return services.
type Services interface {
	Auth() AuthService
}

// NewServices creates a new instance of Services.
func NewServices(db *gorm.DB) Services {
	s := &services{}

	s.User = &userService{
		db: db,
	}

	s.auth = &authService{
		db:          db,
		userService: s.User,
	}

	s.Student = &studentService{
		db: db,
	}

	return s
}

// services contains all services structures in one whole place.
type services struct {
	auth    AuthService
	User    *userService
	Student *studentService
}

func (s *services) Auth() AuthService {
	return s.auth
}
