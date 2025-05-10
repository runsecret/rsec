package secretref

type VaultType int

const (
	VaultTypeAwsSecretsManager    VaultType = iota // AWS
	VaultTypeAzureKeyVault                         // Azure
	VaultTypeAzureKeyVaultChina                    // Azure China
	VaultTypeAzureKeyVaultUSGov                    // Azure US Gov
	VaultTypeAzureKeyVaultGermany                  // Azure Germany
	VaultTypeHashicorpVault                        // HashiCorp Vault
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
	case "kv.hashi":
		return VaultTypeHashicorpVault
	default:
		return VaultTypeUnknown
	}
}
