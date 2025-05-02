package vault

import (
	"errors"
	"regexp"

	"github.com/runsecret/rsec/pkg/aws"
	"github.com/runsecret/rsec/pkg/azure"
	"github.com/runsecret/rsec/pkg/secretref"
)

type Client struct {
	awsClient   *aws.SecretsManager
	azureClient *azure.KeyVault
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
	case secretref.VaultTypeAzureKeyVault:
		if c.azureClient == nil {
			c.azureClient = azure.NewKeyVault()
		}
		secret, err = c.azureClient.GetSecret(secretRef.GetVaultAddress())
	default:
		return "", errors.New("secret vault type unsupported")
	}

	return
}

func GetIdentifierType(secretID string) SecretIdentifierType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)                                                                                                                         // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	azureAddrRegex := regexp.MustCompile(`^https:\/\/(?:(?:[^\/]+\.vault\.(azure\.(net|cn)|usgovcloudapi\.net|microsoftazure\.de))|localhost(?::\d+)?)\/secrets\/.*?\/?$`) // Ex: https://myvaultname.vault.azure.net/secrets/mysecretname/

	switch {
	case secretref.IsSecretRef(secretID):
		return SecretIdentifierTypeRef
	case awsArnRegex.MatchString(secretID):
		return SecretIdentifierTypeAwsArn
	case azureAddrRegex.MatchString(secretID):
		return SecretIdentifierTypeAzureArn
	default:
		return SecretIdentifierTypeUnknown
	}
}
