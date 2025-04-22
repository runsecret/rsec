package secretref

type VaultType int

const (
	VaultTypeAwsSecretsManager VaultType = iota // AWS
	VaultTypeGcpSecretsManager                  // GCP - TODO: Implement
	VaultTypeAzureKeyVault                      // Azure - TODO: Implement
	VaultTypeUnknown                            // Unknown
)

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
