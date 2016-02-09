package config

import (
	"errors"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/ab22/env"
)

// Config struct that contains all of the configuration variables
// that are set up in the environment.
type Config struct {
	DbLogMode       bool
	NoReplyEmail    string
	SessionLifeTime time.Duration

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

// NewConfig initializes a new Config structure.
func NewConfig() *Config {
	var (
		adminAppFolder string
		err            error

		cfg = &Config{
			DbLogMode:       false,
			NoReplyEmail:    "noreply@abcd.com",
			SessionLifeTime: time.Minute * 30,
		}
	)

	if err = env.Parse(cfg); err != nil {
		log.Fatalln(err)
	}

	if cfg.App.Env == "DEV" {
		adminAppFolder = "app"
	} else {
		adminAppFolder = "dist"
	}

	cfg.App.Frontend.Admin = path.Join("frontend/admin", adminAppFolder)

	return cfg
}

// Validate checks if the most important fields are set and are not empty
// values.
func (c *Config) Validate() error {
	var errorMsg = "config: field [%v] was not set!"

	// App validation
	if c.App.HostUrl == "" {
		return errors.New(fmt.Sprintf(errorMsg, "App.HostURL"))
	}

	if c.App.Secret == "" {
		return errors.New(fmt.Sprintf(errorMsg, "App.Secret"))
	}

	//Db validation
	if c.Db.Host == "" {
		return errors.New(fmt.Sprintf(errorMsg, "Db.Host"))
	}

	if c.Db.Port == 0 {
		return errors.New(fmt.Sprintf(errorMsg, "Db.Port"))
	}

	if c.Db.User == "" {
		return errors.New(fmt.Sprintf(errorMsg, "Db.User"))
	}

	if c.Db.Password == "" {
		return errors.New(fmt.Sprintf(errorMsg, "Db.Password"))
	}

	if c.Db.Name == "" {
		return errors.New(fmt.Sprintf(errorMsg, "Db.Name"))
	}

	return nil
}

// Print configuration values to the log. Some user and password fields
// are omitted for security reasons.
func (c *Config) Print() {
	log.Println("----------------------------------")
	log.Println("Application Port:", c.App.Port)
	log.Println("       Admin App:", c.App.Frontend.Admin)
	log.Println("     Environment:", c.App.Env)
	log.Println("   Database Host:", c.Db.Host)
	log.Println("   Database Port:", c.Db.Port)
	log.Println("   Database Name:", c.Db.Name)
	log.Println("----------------------------------")
}
