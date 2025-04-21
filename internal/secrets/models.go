package secrets

type VaultType int

const (
	VaultTypeAws     VaultType = iota // AWS
	VaultTypeGcp                      // GCP - TODO: Implement
	VaultTypeAzure                    // Azure - TODO: Implement
	VaultTypeUnknown                  // Unknown
)

type SecretIdentifierType int

const (
	SecretIdentifierTypeAwsArn   SecretIdentifierType = iota // AWS ARN
	SecretIdentifierTypeAwsRef                               // AWS Ref
	SecretIdentifierTypeAzureArn                             // Azure ARN
	SecretIdentifierTypeAzureRef                             // Azure Ref
	SecretIdentifierTypeUnknown                              // Unknown
)
