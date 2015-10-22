package config

import (
	"path"

	"github.com/ab22/env"
)

// Constant config variables.
const (
	DbLogMode    = false
	NoReplyEmail = "noreply@abcd.es"
)

var (
	EnvVariables envVariables // Contains environment configuration variables.
)

func Initialize() error {
	var adminFolder string
	var err error

	if err = env.Parse(&EnvVariables); err != nil {
		return err
	}

	if EnvVariables.App.Env != "DEV" {
		adminFolder = "dist"
	}

	EnvVariables.App.Frontend.Admin = path.Join("frontend/admin", adminFolder)

	if err = EnvVariables.Validate(); err != nil {
		return err
	}

	return nil
}
