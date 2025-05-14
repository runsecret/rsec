package vault

type SecretIdentifierType int

const (
	SecretIdentifierTypeAwsArn   SecretIdentifierType = iota // AWS ARN
	SecretIdentifierTypeAzureArn                             // Azure ARN
	SecretIdentifierTypeHashiURL                             // HashiCorp Path
	SecretIdentifierTypeRef                                  // rsec Ref
	SecretIdentifierTypeUnknown                              // Unknown
)
