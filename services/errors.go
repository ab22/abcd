package services

import (
	"errors"
)

// Define global errors
var (
	DuplicateUsernameError = errors.New("user service: edit user: username already exists")
)
