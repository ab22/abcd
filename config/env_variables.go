package config

import (
	"errors"
	"fmt"
	"log"
)

// Struct that contains all of the configuration variables
// that are set up in the environment.
type envVariables struct {
	App struct {
		HostUrl string `env:"HOST_URL" envDefault:"http://localhost:1337/"`
		Secret  string `env:"SECRET_KEY" envDefault:"SOME-VERY-SECRET-AND-RANDOM-KEY"`
		Port    int    `env:"PORT" envDefault:"1337"`
		Env     string `env:"ENV" envDefault:"DEV"`

		Frontend struct {
			Admin string
		}
	}
}

// Validate checks if the most important fields are set and are not empty
// values.
func (e *envVariables) Validate() error {
	var errorMsg = "config: field [%v] was not set!"

	if e.App.HostUrl == "" {
		return errors.New(fmt.Sprintf(errorMsg, "App.HostURL"))
	}

	if e.App.Secret == "" {
		return errors.New(fmt.Sprintf(errorMsg, "App.Secret"))
	}

	return nil
}

// Print configuration values to the log. Some user and password fields
// are omitted for security reasons.
func (e *envVariables) Print() {
	log.Println("----------------------------------")
	log.Println("Application Port:", e.App.Port)
	log.Println("       Admin App:", e.App.Frontend.Admin)
	log.Println("     Environment:", e.App.Env)
	log.Println("----------------------------------")
}
