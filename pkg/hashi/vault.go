package hashi

import (
	"context"
	"encoding/json"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
	"github.com/runsecret/rsec/pkg/secretref"
)

type KVClient interface {
	Get(ctx context.Context, path string) (*vault.KVSecret, error)
}

type VaultClientAPI interface {
	KVv2(mountpath string) KVClient
}

type vaultClientWrapper struct {
	client *vault.Client
}

func (v *vaultClientWrapper) KVv2(mountpath string) KVClient {
	return v.client.KVv2(mountpath)
}

type Vault struct {
	client VaultClientAPI
}

func NewVault() *Vault {
	return &Vault{}
}

func (h *Vault) getClient() VaultClientAPI {
	// Create client if not already created
	if h.client == nil {
		// Setup config for HashiCorp Vault
		config := vault.DefaultConfig()
		vaultAddr := os.Getenv("VAULT_ADDR")
		if vaultAddr == "" {
			log.Fatal("VAULT_ADDR environment variable is not set")
		}
		config.Address = vaultAddr

		// Create a new Vault client
		client, err := vault.NewClient(config)
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

func (h *Vault) GetSecret(ref secretref.SecretReference) (string, error) {
	client := h.getClient()

	// Get the secret from the vault
	kvClient := client.KVv2(ref.VaultProviderAddress)
	secret, err := kvClient.Get(context.Background(), ref.SecretName)
	if err != nil {
		return "", err
	}

	// Convert secret map to JSON string
	jsonBytes, err := json.Marshal(secret.Data)
	if err != nil {
		return "", err
	}
	secretJsonString := string(jsonBytes)

	return secretJsonString, nil
}
