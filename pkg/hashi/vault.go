package hashi

import (
	"context"
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
)

type HashiKVClient interface {
	Get(ctx context.Context, path string) (*vault.KVSecret, error)
}

type HashiVaultClientAPI interface {
	KVv2(mountpath string) HashiKVClient
}

type vaultClientWrapper struct {
	client *vault.Client
}

func (v *vaultClientWrapper) KVv2(mountpath string) HashiKVClient {
	return v.client.KVv2(mountpath)
}

type HashiVault struct {
	client HashiVaultClientAPI
}

func NewHashiVault() *HashiVault {
	return &HashiVault{}
}

func (h *HashiVault) getClient(vaultAddress string) HashiVaultClientAPI {
	// Create client if not already created
	if h.client == nil {
		client, err := vault.NewClient(vault.DefaultConfig())
		if err != nil {
			log.Fatalf("Cannot connect to HashiCorp Vault client: %s", err)
		}

		// TODO: Replace with real auth
		client.SetToken("dev-only-token")

		h.client = &vaultClientWrapper{client: client}
	}

	// Return the client
	return h.client
}

func (h *HashiVault) GetSecret(vaultAddress string, secretPath string) (string, error) {
	client := h.getClient(vaultAddress)

	// Get the secret from the vault
	kvClient := client.KVv2("secret")
	secret, err := kvClient.Get(context.Background(), "my-project")
	if err != nil {
		return "", err
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		return "", fmt.Errorf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	return value, nil
}
