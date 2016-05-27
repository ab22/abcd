package models

// Role is a simple integer.
type Role int

const (
	// Admin role.
	Admin Role = iota
	// Teacher role.
	Teacher
)
