package services

import (
	"errors"
)

// Define global errors
var (
	DuplicateUsernameError        = errors.New("username already exists")
	DuplicateStudentIdNumberError = errors.New("id number/passport already exists")
	RecordNotFound                = errors.New("record not found")
)
