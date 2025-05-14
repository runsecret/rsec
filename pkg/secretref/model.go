package secretref

type VaultType int

const (
	VaultTypeAwsSecretsManager    VaultType = iota // AWS
	VaultTypeAzureKeyVault                         // Azure
	VaultTypeAzureKeyVaultChina                    // Azure China
	VaultTypeAzureKeyVaultUSGov                    // Azure US Gov
	VaultTypeAzureKeyVaultGermany                  // Azure Germany
	VaultTypeHashicorpVaultKv1                     // HashiCorp Kv1 Vault Engine Secret
	VaultTypeHashicorpVaultKv2                     // HashiCorp Kv2 Vault Engine Secret
	VaultTypeHashicorpVaultCred                    // HashiCorp Vault Credential Secret
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
	case "kv1.hashi":
		return VaultTypeHashicorpVaultKv1
	case "kv2.hashi":
		return VaultTypeHashicorpVaultKv2
	case "cred.hashi":
		return VaultTypeHashicorpVaultCred
	default:
		return VaultTypeUnknown
	}
}
