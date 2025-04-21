package secrets

import (
	"regexp"

	"github.com/runsecret/rsec/pkg/aws"
	"github.com/runsecret/rsec/pkg/azure"
)

type vaultClient interface {
	GetSecret(secretRef string) (string, error)
}

type VaultClient struct {
	awsClient   vaultClient
	azureClient vaultClient
}

func NewVaultClient() VaultClient {
	return VaultClient{}
}

func (vc VaultClient) CheckForSecret(secretRef string) (secret string, err error) {
	vaultType, vaultAddress := GetVaultAddress(secretRef)

	switch vaultType {
	case VaultTypeAws:
		if vc.awsClient == nil {
			vc.awsClient = aws.NewSecretsManager()
		}
		secret, err = vc.awsClient.GetSecret(vaultAddress)
	case VaultTypeAzure:
		if vc.azureClient == nil {
			vc.azureClient = azure.NewKeyVault()
		}
		secret, err = vc.azureClient.GetSecret(vaultAddress)
	default:
		// Do nothing
	}

	return
}

func GetVaultAddress(secretRef string) (VaultType, string) {
	switch GetIdentifierType(secretRef) {
	case SecretIdentifierTypeAwsRef:
		return VaultTypeAws, ConvertAwsRefToAwsArn(secretRef)
	case SecretIdentifierTypeAzureRef:
		return VaultTypeAzure, ConvertAzureRefToAzureArn(secretRef)
	default:
		return VaultTypeUnknown, "Invalid secret reference"
	}
}

func GetIdentifierType(secretRef string) SecretIdentifierType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)                                        // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	awsRefRegex := regexp.MustCompile(`aws:\/\/[^\/]*\/[^\/]*\/[^\/]*`)                   // Ex: aws://us-west-2/123456789012/my-secret
	azureArnRegex := regexp.MustCompile(`https:\/\/.*\.vault\.azure\.net\/secrets\/.*\/`) // Ex: https://myvaultname.vault.azure.net/secrets/mysecretname/
	azureRefRegex := regexp.MustCompile(`azure:\/\/[^\/]*\/[^\/]*`)                       // Ex: azure://myvaultname/mysecretname

	switch {
	case awsArnRegex.MatchString(secretRef):
		return SecretIdentifierTypeAwsArn
	case awsRefRegex.MatchString(secretRef):
		return SecretIdentifierTypeAwsRef
	case azureArnRegex.MatchString(secretRef):
		return SecretIdentifierTypeAzureArn
	case azureRefRegex.MatchString(secretRef):
		return SecretIdentifierTypeAzureRef
	default:
		return SecretIdentifierTypeUnknown
	}
}
