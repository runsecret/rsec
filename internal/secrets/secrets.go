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

func (vc VaultClient) CheckForSecret(secretUrl string) (secret string, err error) {
	secretRef, err := NewSecretReferenceFromURL(secretUrl)
	if err != nil {
		return "", err
	}

	switch secretRef.vaultType {
	case VaultTypeAwsSecretsManager:
		if vc.awsClient == nil {
			vc.awsClient = aws.NewSecretsManager()
		}
		secret, err = vc.awsClient.GetSecret(secretRef.GetVaultAddress())
	default:
		// Do nothing
	}

	return
}

func GetIdentifierType(secretID string) SecretIdentifierType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`) // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	refRegex := regexp.MustCompile(`rsec://.*`)    // Ex: aws://us-west-2/123456789012/my-secret

	switch {
	case awsArnRegex.MatchString(secretID):
		return SecretIdentifierTypeAwsArn
	case refRegex.MatchString(secretID):
		return SecretIdentifierTypeRef
	default:
		return SecretIdentifierTypeUnknown
	}
}
