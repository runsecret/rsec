package hashi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
	"github.com/runsecret/rsec/pkg/secretref"
)

type KVClient interface {
	Get(ctx context.Context, path string) (*vault.KVSecret, error)
}

type LogicalClient interface {
	Read(path string) (*vault.Secret, error)
}

type VaultClientAPI interface {
	KVv1(mountpath string) KVClient
	KVv2(mountpath string) KVClient
	Logical() LogicalClient
}

type Vault struct {
	client VaultClientAPI
}

func NewVault() *Vault {
	return &Vault{}
}

func (h *Vault) getClient(ref secretref.SecretReference) VaultClientAPI {
	// Create client if not already created
	if h.client == nil {
		// Setup config for HashiCorp Vault
		config := vault.DefaultConfig()
		vaultAddr := ref.Endpoint
		if vaultAddr == "" {
			vaultAddr = os.Getenv("VAULT_ADDR")
		}
		if vaultAddr == "" {
			log.Fatal("VAULT_ADDR environment variable is not set")
		}
		config.Address = vaultAddr

		// Create a new Vault client
		client, err := vault.NewClient(config)
		if err != nil {
			log.Fatalf("Cannot connect to HashiCorp Vault client: %s", err)
		}

		err = authenticateWithToken(client)
		if err != nil {
			log.Fatalf("Failed to authenticate with Vault: %s", err)
		}

		h.client = &vaultClientWrapper{client: client}
	}

	// Return the client
	return h.client
}

func (h *Vault) GetKv1Secret(ref secretref.SecretReference) (string, error) {
	client := h.getClient(ref)

	// Get the secret from the vault
	kvClient := client.KVv1(ref.VaultProviderAddress)
	secret, err := kvClient.Get(context.Background(), ref.SecretName)
	if err != nil {
		return "", err
	}

	// Convert secret map to JSON string
	jsonBytes, err := json.Marshal(secret.Data)
	if err != nil {
		return "", err
	}
	secretJSONString := string(jsonBytes)

	return secretJSONString, nil
}

func (h *Vault) GetCredential(ref secretref.SecretReference) (string, error) {
	client := h.getClient(ref)

	// Get the secret from the vault
	secretPath := fmt.Sprintf("%s/creds/%s", ref.VaultProviderAddress, ref.SecretName)
	secret, err := client.Logical().Read(secretPath)
	if err != nil {
		return "", err
	}

	// Convert secret map to JSON string
	jsonBytes, err := json.Marshal(secret.Data)
	if err != nil {
		return "", err
	}
	secretJSONString := string(jsonBytes)

	return secretJSONString, nil
}

// GetKv2Secret retrieves a secret from HashiCorp Vault using the KV v2 API
func (h *Vault) GetKv2Secret(ref secretref.SecretReference) (string, error) {
	client := h.getClient(ref)

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
	secretJSONString := string(jsonBytes)

	return secretJSONString, nil
}
