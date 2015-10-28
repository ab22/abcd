# env

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
