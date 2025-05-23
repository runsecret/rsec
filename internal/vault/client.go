package vault

import (
	"errors"
	"regexp"

	"github.com/runsecret/rsec/pkg/aws"
	"github.com/runsecret/rsec/pkg/azure"
	"github.com/runsecret/rsec/pkg/hashi"
	"github.com/runsecret/rsec/pkg/secretref"
)

type Client struct {
	awsClient   *aws.SecretsManager
	azureClient *azure.KeyVault
	hashiClient *hashi.Vault
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
		secret, err = c.azureClient.GetSecret(secretRef)
	case secretref.VaultTypeHashicorpVaultKv1:
		if c.hashiClient == nil {
			c.hashiClient = hashi.NewVault()
		}
		secret, err = c.hashiClient.GetKv1Secret(secretRef)
	case secretref.VaultTypeHashicorpVaultKv2:
		if c.hashiClient == nil {
			c.hashiClient = hashi.NewVault()
		}
		secret, err = c.hashiClient.GetKv2Secret(secretRef)
	case secretref.VaultTypeHashicorpVaultCred:
		if c.hashiClient == nil {
			c.hashiClient = hashi.NewVault()
		}
		secret, err = c.hashiClient.GetCredential(secretRef)
	default:
		return "", errors.New("secret vault type unsupported")
	}

	return
}

func GetIdentifierType(secretID string) SecretIdentifierType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)                                                                                                      // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	azureAddrRegex := regexp.MustCompile(`^https:\/\/(?:(?:[^\/]+\.vault\.(azure\.(net|cn)|usgovcloudapi\.net|microsoftazure\.de)))\/secrets\/.*?\/?$`) // Ex: https://myvaultname.vault.azure.net/secrets/mysecretname/
	hashiURLRegex := regexp.MustCompile(`https?:\/\/(([-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6})|(localhost|(\d{1,3}\.){3}\d{1,3}))(:[0-9]{1,5})?\b([-a-zA-Z0-9()@:%_\+.~#?&\/=]*)\/v1\/[^\/]+\/(data|creds)\/.*`)

	switch {
	case secretref.IsSecretRef(secretID):
		return SecretIdentifierTypeRef
	case awsArnRegex.MatchString(secretID):
		return SecretIdentifierTypeAwsArn
	case azureAddrRegex.MatchString(secretID):
		return SecretIdentifierTypeAzureArn
	case hashiURLRegex.MatchString(secretID):
		return SecretIdentifierTypeHashiURL
	default:
		return SecretIdentifierTypeUnknown
	}
}
