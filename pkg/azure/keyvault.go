package azure

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

// Define the interface to match the Azure Key Vault client methods
type KeyVaultClientAPI interface {
	GetSecret(ctx context.Context, secretName string, version string, opt *azsecrets.GetSecretOptions) (azsecrets.GetSecretResponse, error)
}

type KeyVault struct {
	client KeyVaultClientAPI
}

func NewKeyVault() KeyVault {
	return KeyVault{}
}

func (k *KeyVault) getClient() KeyVaultClientAPI {
	vaultURI := fmt.Sprintf("https://%s.vault.azure.net/", os.Getenv("KEY_VAULT_NAME"))

	// Create client if not already created
	if k.client == nil {
		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			log.Fatalf("Cannot load the Azure credentials: %s", err)
		}

		// Establish a connection to the Key Vault client
		client, err := azsecrets.NewClient(vaultURI, cred, nil)
		if err != nil {
			log.Fatalf("Cannot connect to azure KeyVault client: %s", err)
		}
		k.client = client
	}

	// Return the client
	return k.client
}

// Ex: https://myvaultname.vault.azure.net/secrets/mysecretname/version123
func (k KeyVault) GetSecret(secretName string) (string, error) {
	// Call the GetSecretValue API
	version := "" // An empty string version gets the latest version of the secret.
	resp, err := k.getClient().GetSecret(context.TODO(), secretName, version, nil)
	if err != nil {
		return "", err
	}

	return *resp.Value, nil
}
