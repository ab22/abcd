package auth

import "github.com/ab22/abcd/models"

// Service interface describes all functions that must be implemented.
type Service interface {
	BasicAuth(string, string) (*models.User, error)
}
