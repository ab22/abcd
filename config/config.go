package config

import (
	"path"

	"github.com/ab22/env"
)

// Constant config variables.
const (
	DbLogMode    = false
	NoReplyEmail = "noreply@abcd.com"
)

var (
	// Contains environment configuration variables.
	EnvVariables envVariables
)

func Initialize() error {
	var adminAppFolder string
	var err error

	if err = env.Parse(&EnvVariables); err != nil {
		return err
	}

	if EnvVariables.App.Env == "DEV" {
		adminAppFolder = "app"
	} else {
		adminAppFolder = "dist"
	}

	EnvVariables.App.Frontend.Admin = path.Join("frontend/admin", adminAppFolder)

	if err = EnvVariables.Validate(); err != nil {
		return err
	}

	return nil
}
