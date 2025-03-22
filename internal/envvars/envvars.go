package envvars

import (
	"fmt"
	"strings"

	"github.com/devenjarvis/signet/internal/aws"
	"github.com/devenjarvis/signet/internal/secretref"
)

func GetSecret(secretRef string) (secret string, err error) {
	vaultType := secretref.ParseVaultType(secretRef)

	switch vaultType {
	case secretref.VaultTypeAws:
		secret, err = aws.GetSecret(secretRef)
	default:
		// Do nothing
	}

	return
}

func SetSecrets(rawEnv []string) (envVars []string, redactList []string, err error) {
	envVars = rawEnv
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
