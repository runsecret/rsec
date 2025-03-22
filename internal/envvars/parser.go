package envvars

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/devenjarvis/signet/internal/aws"
)

func ParseVaultType(secretRef string) VaultType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)

	switch {
	case awsArnRegex.MatchString(secretRef):
		return VaultTypeAws
	default:
		return VaultTypeUnknown
	}
}

func SetSecrets(rawEnv []string) (envVars []string, redactList []string, err error) {
	envVars = rawEnv
	for i, envVar := range envVars {
		parts := strings.SplitN(envVar, "=", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		var secret string
		switch ParseVaultType(value) {
		case VaultTypeAws:
			secret, err = aws.GetSecret(value)
		case VaultTypeUnknown:
			// Leave it alone
			continue
		}

		if err != nil {
			return
		}

		// Replace the secret in the env var
		envVars[i] = fmt.Sprintf("%s=%s", key, secret)

		// Add secret to list of secrets for redaction
		redactList = append(redactList, secret)
	}

	return
}
