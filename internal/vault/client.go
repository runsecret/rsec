package vault

import (
	"regexp"

	"github.com/runsecret/rsec/pkg/aws"
	"github.com/runsecret/rsec/pkg/secretref"
)

type Client struct {
	awsClient *aws.SecretsManager
}

func NewClient() Client {
	return Client{}
}

func (c Client) GetSecret(secretID string) (secret string, err error) {
	secretRef, err := secretref.NewFromString(secretID)
	if err != nil {
		return "", err
	}

	switch secretRef.VaultType {
	case secretref.VaultTypeAwsSecretsManager:
		if c.awsClient == nil {
			c.awsClient = aws.NewSecretsManager()
		}
		secret, err = c.awsClient.GetSecret(secretRef.GetVaultAddress())
	default:
		// TODO: Return error if the vault type is not supported
	}

	return
}

func GetIdentifierType(secretID string) SecretIdentifierType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`) // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret

	switch {
	case secretref.IsSecretRef(secretID):
		return SecretIdentifierTypeRef
	case awsArnRegex.MatchString(secretID):
		return SecretIdentifierTypeAwsArn
	default:
		return SecretIdentifierTypeUnknown
	}
}
