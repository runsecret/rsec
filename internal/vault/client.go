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

func (c Client) CheckForSecret(secretID string) (secretValue string, err error) {
	secretRef, err := secretref.NewFromString(secretID)
	if err != nil {
		return "", err
	}

	switch secretRef.VaultType {
	case secretref.VaultTypeAwsSecretsManager:
		if c.awsClient == nil {
			c.awsClient = aws.NewSecretsManager()
		}
		secretValue, err = c.awsClient.GetSecret(secretRef.GetVaultAddress())
	default:
		// Do nothing
	}

	return
}

func GetIdentifierType(secretID string) SecretIdentifierType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)                                                                   // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	refRegex := regexp.MustCompile(`(rsec:\/\/)([-a-zA-Z0-9_\+~#=]*)\.([a-z.]*)\/([-a-zA-Z0-9()@:%_\+.~#?&\/\/=]*)`) // Ex: rsec://123456789012.aws/my-secret?region=us-west-2

	switch {
	case awsArnRegex.MatchString(secretID):
		return SecretIdentifierTypeAwsArn
	case refRegex.MatchString(secretID):
		return SecretIdentifierTypeRef
	default:
		return SecretIdentifierTypeUnknown
	}
}
