package secrets

import (
	"regexp"

	"github.com/runsecret/rsec/pkg/aws"
)

func GetVaultAddress(secretRef string) (VaultType, string) {
	switch GetRefType(secretRef) {
	case SecretRefTypeAwsRef:
		return VaultTypeAws, ConvertAwsRefToAwsArn(secretRef)
	default:
		return VaultTypeUnknown, "Invalid secret reference"
	}
}

func GetRefType(secretRef string) SecretRefType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)                      // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	awsRefRegex := regexp.MustCompile(`aws:\/\/[^\/]*\/[^\/]*\/[^\/]*`) // Ex: aws://us-west-2/123456789012/my-secret

	switch {
	case awsArnRegex.MatchString(secretRef):
		return SecretRefTypeAwsArn
	case awsRefRegex.MatchString(secretRef):
		return SecretRefTypeAwsRef
	default:
		return SecretRefTypeUnknown
	}
}

func GetSecret(secretRef string) (secret string, err error) {
	vaultType, vaultAddress := GetVaultAddress(secretRef)

	switch vaultType {
	case VaultTypeAws:
		secret, err = aws.GetSecret(vaultAddress)
	default:
		// Do nothing
	}

	return
}
