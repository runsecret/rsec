package vault

type SecretIdentifierType int

const (
	SecretIdentifierTypeAwsArn   SecretIdentifierType = iota // AWS ARN
	SecretIdentifierTypeAzureArn                             // Azure ARN
	SecretIdentifierTypeRef                                  // rsec Ref
	SecretIdentifierTypeUnknown                              // Unknown
)
