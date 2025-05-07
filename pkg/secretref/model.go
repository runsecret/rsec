package secretref

type VaultType int

const (
	VaultTypeAwsSecretsManager    VaultType = iota // AWS
	VaultTypeGcpSecretsManager                     // GCP - TODO: Implement
	VaultTypeAzureKeyVault                         // Azure
	VaultTypeAzureKeyVaultChina                    // Azure China
	VaultTypeAzureKeyVaultUSGov                    // Azure US Gov
	VaultTypeAzureKeyVaultGermany                  // Azure Germany
	VaultTypeUnknown                               // Unknown
)

func vaultTypeFromString(vaultType string) VaultType {
	switch vaultType {
	case "sm.aws":
		return VaultTypeAwsSecretsManager
	case "kv.azure":
		return VaultTypeAzureKeyVault
	case "kv.azure.cn":
		return VaultTypeAzureKeyVaultChina
	case "kv.azure.us":
		return VaultTypeAzureKeyVaultUSGov
	case "kv.azure.de":
		return VaultTypeAzureKeyVaultGermany
	case "sm.gcp":
		return VaultTypeGcpSecretsManager
	default:
		return VaultTypeUnknown
	}
}
