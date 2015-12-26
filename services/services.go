package services

import (
	"github.com/jinzhu/gorm"
)

type Services struct {
	Auth    *authService
	User    *userService
	Student *studentService
}

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
