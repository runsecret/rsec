package hashi

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	vault "github.com/hashicorp/vault/api"
	"github.com/mitchellh/go-homedir"
)

// authenticateWithToken authenticates using a token
func authenticateWithToken(client *vault.Client) error {
	var token string

	// Try to get token from environment variable
	token = os.Getenv("VAULT_TOKEN")
	if token != "" {
		client.SetToken(token)
		return nil
	}

	// Try to read token from default token file
	home, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	tokenPath := filepath.Join(home, ".vault-token")

	// Check if token file exists
	if _, err := os.Stat(tokenPath); os.IsNotExist(err) {
		return errors.New("no token found in ~/.vault-token or VAULT_TOKEN environment variable")
	}

	// Read token from file
	tokenBytes, err := os.ReadFile(tokenPath)
	if err != nil {
		return fmt.Errorf("failed to read token file: %w", err)
	}

	token = strings.TrimSpace(string(tokenBytes))
	client.SetToken(token)
	return nil
}
