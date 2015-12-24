# env

[![Build Status](https://travis-ci.org/ab22/env.svg)](https://travis-ci.org/ab22/env)
[![GoDoc](https://godoc.org/github.com/ab22/env?status.svg)](https://godoc.org/github.com/ab22/env)

Use tag structures to parse environment variables into structure fields.

Having structs inside structs parsed is now possible. Updated the example below with fully working code.

## Example

```go
package main

import (
	"github.com/ab22/env"
	"log"
)

type AppConfig struct {
	Env        string `env:"ENV" envDefault:"PRODUCTION"`
	RiotApiKey string `env:"RIOT_API_KEY"`
	Port       int    `env:"APP_PORT" envDefault:"1337"`

	Smtp struct {
		Host     string `env:"SMTP_HOST"`
		Port     int    `env:"SMTP_PORT"`
		User     string `env:"SMTP_USER"`
		Password string `env:"SMTP_PASS"`
	}

	Db struct {
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		Port     int    `env:"DB_PORT" envDefault:"5432"`
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASS" envDefault:"1234"`
		Name     string `env:"DB_NAME" envDefault:"lol_db"`
	}
}

// Print configuration values to the log. Some user and password fields
// are omitted for security reasons.
func (c *AppConfig) Print() {
	log.Println("----------------------------------")
	log.Println("Application Port:", c.Port)
	log.Println("     Environment:", c.Env)
	log.Println("       SMTP Host:", c.Smtp.Host)
	log.Println("       SMTP User:", c.Smtp.User)
	log.Println("       SMTP Port:", c.Smtp.Port)
	log.Println("   Database Host:", c.Db.Host)
	log.Println("   Database Port:", c.Db.Port)
	log.Println("   Database Name:", c.Db.Name)
	log.Println("----------------------------------")
}

func main() {
	config := &AppConfig{}

	env.Parse(config)
	config.Print()
}
```
The example above prints:
```
2015/11/01 18:01:13 ----------------------------------
2015/11/01 18:01:13 Application Port: 1337
2015/11/01 18:01:13      Environment: PRODUCTION
2015/11/01 18:01:13        SMTP Host: smtp.mandrillapp.com
2015/11/01 18:01:13        SMTP User: app32793597@heroku.com
2015/11/01 18:01:13        SMTP Port: 587
2015/11/01 18:01:13    Database Host: localhost
2015/11/01 18:01:13    Database Port: 5432
2015/11/01 18:01:13    Database Name: lol_db
2015/11/01 18:01:13 ----------------------------------
```
