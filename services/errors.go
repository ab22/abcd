package services

import (
	"errors"
)

// Define global errors
var (
	ErrDuplicatedUsername        = errors.New("username already exists")
	ErrDuplicatedStudentIDNumber = errors.New("id number/passport already exists")
	ErrRecordNotFound            = errors.New("record not found")
)
