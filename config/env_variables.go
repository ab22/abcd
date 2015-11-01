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

	Db struct {
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		Port     int    `env:"DB_PORT" envDefault:"5432"`
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASS" envDefault:"1234"`
		Name     string `env:"DB_NAME" envDefault:"abcd"`
	}
}

// Validate checks if the most important fields are set and are not empty
// values.
func (e *envVariables) Validate() error {
	var errorMsg = "config: field [%v] was not set!"

	// App validation
	if e.App.HostUrl == "" {
		return errors.New(fmt.Sprintf(errorMsg, "App.HostURL"))
	}

	if e.App.Secret == "" {
		return errors.New(fmt.Sprintf(errorMsg, "App.Secret"))
	}

	//Db validation
	if e.Db.Host == "" {
		return errors.New(fmt.Sprintf(errorMsg, "Db.Host"))
	}

	if e.Db.Port == 0 {
		return errors.New(fmt.Sprintf(errorMsg, "Db.Port"))
	}

	if e.Db.User == "" {
		return errors.New(fmt.Sprintf(errorMsg, "Db.User"))
	}

	if e.Db.Password == "" {
		return errors.New(fmt.Sprintf(errorMsg, "Db.Password"))
	}

	if e.Db.Name == "" {
		return errors.New(fmt.Sprintf(errorMsg, "Db.Name"))
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
	log.Println("   Database Host:", e.Db.Host)
	log.Println("   Database Port:", e.Db.Port)
	log.Println("   Database Name:", e.Db.Name)
	log.Println("----------------------------------")
}
