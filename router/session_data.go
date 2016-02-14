package router

import "time"

const (
	// SessionCookieName is the name used to read the session cookie from the
	// client.
	SessionCookieName = "_session"
)

// SessionData describes the session cookie for all users.
type SessionData struct {
	UserID    int
	RoleID    int
	Email     string
	IsAdmin   bool
	IsTeacher bool
	ExpiresAt time.Time
}
