package envfile

import (
	"os"
	"slices"
	"strings"
)

var ErrInvalidEnvFile = os.ErrInvalid

func ReadEnvFile(path string) (envVars []string, err error) {
	envs, err := os.ReadFile(path)
	if err != nil {
		return
	}
	// Split the env file into individual lines
	envVars = strings.Split(string(envs), "\n")

	// Remove any empty lines or comments
	for i := 0; i < len(envVars); {
		envVar := envVars[i]
		if strings.TrimSpace(envVar) == "" || strings.HasPrefix(envVar, "#") {
			envVars = slices.Delete(envVars, i, i+1)
		} else {
			i++
		}
	}

	err = validateEnvFile(envVars)
	if err != nil {
		envVars = []string{}
		return
	}

	return
}

func validateEnvFile(envVars []string) (err error) {
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
