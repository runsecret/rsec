package envfile

import (
	"os"
	"strings"
)

var ErrInvalidEnvFile = os.ErrInvalid

func validate(envVars []string) (err error) {
	// Valiate the env file is in the correct format
	for _, envVar := range envVars {
		parts := strings.SplitN(envVar, "=", 2)
		// If there is not exactly one equals sign, raise an error
		if len(parts) != 2 {
			err = ErrInvalidEnvFile
			return
		}
		// If the key is empty, raise an error
		if strings.TrimSpace(parts[0]) == "" {
			err = ErrInvalidEnvFile
			return
		}
		// If the value is empty, raise an error
		if strings.TrimSpace(parts[1]) == "" {
			err = ErrInvalidEnvFile
			return
		}
	}

	return
}

func Read(path string) (envVars []string, err error) {
	envs, err := os.ReadFile(path)
	if err != nil {
		return
	}
	envVars = strings.Split(string(envs), "\n")

	err = validate(envVars)
	if err != nil {
		envVars = []string{}
		return
	}

	return
}
