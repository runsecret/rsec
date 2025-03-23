package envvars

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/runsecret/rsec/internal/secrets"
	"github.com/runsecret/rsec/pkg/aws"
)

func GetSecret(secretRef string) (secret string, err error) {
	vaultType, vaultRef := secrets.GetVaultReference(secretRef)

	switch vaultType {
	case secrets.VaultTypeAws:
		secret, err = aws.GetSecret(vaultRef)
	default:
		// Do nothing
	}

	return
}

func SetSecrets(cmd *exec.Cmd, envFilePath string) (envVars []string, redactList []string, err error) {
	// load ENV VARs
	envVars, err = loadEnvVars(cmd, envFilePath)
	if err != nil {
		return
	}

	// Replace secret references in ENV VARS
	for i, envVar := range envVars {
		// Split env vars
		parts := strings.SplitN(envVar, "=", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Try to get secret from env var
		var secret string
		secret, err = GetSecret(value)
		if err != nil {
			return
		}

		// If secret was found, replace it in the env var
		if secret != "" {
			// Replace the secret in the env var
			envVars[i] = fmt.Sprintf("%s=%s", key, secret)
			// Add secret to list of secrets for redaction
			redactList = append(redactList, secret)
		}
	}

	return
}

func loadEnvVars(cmd *exec.Cmd, envFilePath string) ([]string, error) {
	// Load system env vars
	cmdEnviron := cmd.Environ()

	// If --env flag used, load env vars from file
	if envFilePath != "" {
		fileEnviron, err := readEnvFile(envFilePath)
		if err != nil {
			return cmdEnviron, err
		}
		cmdEnviron = append(cmdEnviron, fileEnviron...)
	}

	return cmdEnviron, nil
}
