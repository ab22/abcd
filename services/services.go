package services

import (
	"github.com/jinzhu/gorm"
)

type Services struct {
	Auth *authService
	User *userService
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

	return s
}
