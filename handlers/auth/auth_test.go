package auth_test

import (
	"testing"

	"github.com/ab22/abcd/handlers/auth"
)

func TestCheckAuth(t *testing.T) {
	// In case we do something else in the future, add tests here.
	var (
		authHandler = auth.NewHandler(nil)
	)

	err := authHandler.CheckAuth(nil, nil, nil)

	if err != nil {
		t.Fatal("error checking authentication:", err)
	}
}

func TestLogin(t *testing.T) {

}
