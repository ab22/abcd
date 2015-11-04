package services

import (
	"errors"
)

// Define global errors
var (
	DuplicateUsernameError = errors.New("username already exists")
	RecordNotFound         = errors.New("record not found")
)
