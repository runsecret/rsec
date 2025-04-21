package secrets

type VaultType int

const (
	VaultTypeAwsSecretsManager VaultType = iota // AWS
	VaultTypeGcpSecretsManager                  // GCP - TODO: Implement
	VaultTypeAzureKeyVault                      // Azure - TODO: Implement
	VaultTypeUnknown                            // Unknown
)

type SecretIdentifierType int

const (
	SecretIdentifierTypeAwsArn  SecretIdentifierType = iota // AWS ARN
	SecretIdentifierTypeRef                                 // rsec Ref
	SecretIdentifierTypeUnknown                             // Unknown
)
