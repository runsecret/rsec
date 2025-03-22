package secrets

type VaultType int

const (
	VaultTypeAws     VaultType = iota // AWS
	VaultTypeGcp                      // GCP - TODO: Implement
	VaultTypeAzure                    // Azure - TODO: Implement
	VaultTypeUnknown                  // Unknown
)

type SecretRefType int

const (
	SecretRefTypeAwsArn  SecretRefType = iota // AWS ARN
	SecretRefTypeAwsRef                       // AWS Ref
	SecretRefTypeUnknown                      // Unknown
)
