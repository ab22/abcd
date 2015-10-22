package config

// Constant config variables.
const (
	DbLogMode    = false
	NoReplyEmail = "noreply@abcd.es"
)

var (
	EnvVariables envVariables // Contains environment configuration variables.
)

func Initialize() error {
	var err error

	if err = EnvVariables.Parse(); err != nil {
		return err
	}

	if err = EnvVariables.Validate(); err != nil {
		return err
	}

	return nil
}
