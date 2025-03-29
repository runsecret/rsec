package secrets

import (
	"regexp"

	"github.com/runsecret/rsec/pkg/aws"
)

type VaultClient struct {
	awsClient *aws.SecretsManager
}

func NewVaultClient() VaultClient {
	return VaultClient{}
}

func (vc VaultClient) CheckForSecret(secretRef string) (secret string, err error) {
	vaultType, vaultAddress := GetVaultAddress(secretRef)

	switch vaultType {
	case VaultTypeAws:
		if vc.awsClient == nil {
			vc.awsClient = aws.NewSecretsManager("")
		}
		secret, err = vc.awsClient.GetSecret(vaultAddress)
	default:
		// Do nothing
	}

	return
}

func GetVaultAddress(secretRef string) (VaultType, string) {
	switch GetIdentifierType(secretRef) {
	case SecretIdentifierTypeAwsRef:
		return VaultTypeAws, ConvertAwsRefToAwsArn(secretRef)
	default:
		return VaultTypeUnknown, "Invalid secret reference"
	}
}

func GetIdentifierType(secretRef string) SecretIdentifierType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)                      // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	awsRefRegex := regexp.MustCompile(`aws:\/\/[^\/]*\/[^\/]*\/[^\/]*`) // Ex: aws://us-west-2/123456789012/my-secret

	switch {
	case awsArnRegex.MatchString(secretRef):
		return SecretIdentifierTypeAwsArn
	case awsRefRegex.MatchString(secretRef):
		return SecretIdentifierTypeAwsRef
	default:
		return SecretIdentifierTypeUnknown
	}
}
