package router

import "time"

const (
	SessionCookieName = "_session"
)

// SessionData describes the session cookie for all users.
type SessionData struct {
	UserId    int
	RoleId    int
	Email     string
	IsAdmin   bool
	IsTeacher bool
	ExpiresAt time.Time
}
