package user

// Status defines an int type to define statuses for the User model.
type Status int

// Defines all user statuses.
const (
	Enabled Status = iota
	Disabled
)
