package secrets

import (
	"strings"
)

func ConvertAwsArnToRef(arn string) string {
	// From: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	// To: rsec://123456789012.sm.aws/v1/my-secret?region=us-west-2
	parts := strings.Split(arn, ":")
	region := parts[3]
	account := parts[4]
	secretName := parts[6]
	secretRef := NewSecretReference(account, VaultTypeAwsSecretsManager, secretName)
	secretRef.SetRegion(region)

	return secretRef.String()
}

func vaultTypeFromString(vaultType string) VaultType {
	switch vaultType {
	case "sm.aws":
		return VaultTypeAwsSecretsManager
	case "kv.azure":
		return VaultTypeAzureKeyVault
	case "sm.gcp":
		return VaultTypeGcpSecretsManager
	default:
		return VaultTypeUnknown
	}
}
