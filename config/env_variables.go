package config

import (
	"errors"
	"fmt"
	"log"
	"path"

	"github.com/ab22/env"
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

// Reads all variables from envinronment
func (e *envVariables) Parse() error {
	var err error
	var distFolder string

	if err = env.Parse(&e.App); err != nil {
		return err
	}

	if e.App.Env != "DEV" {
		distFolder = "dist"
	}

	e.App.Frontend.Admin = path.Join("frontend/admin", distFolder)

	return nil
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
	log.Println(" Frontend folder:", e.App.Frontend.Admin)
	log.Println("     Environment:", e.App.Env)
	log.Println("----------------------------------")
}
