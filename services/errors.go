package services

import (
	"errors"
)

// Define global errors
var (
	DuplicatedUsernameError        = errors.New("username already exists")
	DuplicatedStudentIdNumberError = errors.New("id number/passport already exists")
	RecordNotFound                 = errors.New("record not found")
)
