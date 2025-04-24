package vault

type SecretIdentifierType int

const (
	SecretIdentifierTypeAwsArn  SecretIdentifierType = iota // AWS ARN
	SecretIdentifierTypeRef                                 // rsec Ref
	SecretIdentifierTypeUnknown                             // Unknown
)
