package services

import (
	"github.com/jinzhu/gorm"
)

// Services interface defines all methods that return services.
type Services interface {
	Auth() AuthService
	User() UserService
	Student() StudentService
}

// NewServices creates a new instance of Services.
func NewServices(db *gorm.DB) Services {
	s := &services{}

	s.user = &userService{
		db: db,
	}

	s.auth = &authService{
		db:          db,
		userService: s.User,
	}

	s.utudent = &studentService{
		db: db,
	}

	return s
}

// services contains all services structures in one whole place.
type services struct {
	auth    AuthService
	user    UserService
	student StudentService
}

func (s *services) Auth() AuthService {
	return s.auth
}
