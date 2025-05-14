package hashi

import (
	"errors"
	"testing"

	"github.com/runsecret/rsec/pkg/secretref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashiVault_GetKv1Secret_Success(t *testing.T) {
	mockData := map[string]interface{}{
		"password": "test-password",
	}
	mockKVClient := &mockHashiKVClient{mockData: mockData, mockErr: nil}
	mockVaultClient := &mockHashiVaultClientAPI{mockKVClient: mockKVClient}

	hv := &Vault{client: mockVaultClient}

	// Setup secret reference
	secretRef := secretref.New("vault", secretref.VaultTypeHashicorpVaultKv2, "secret/my-project")

	secret, err := hv.GetKv1Secret(secretRef)
	require.NoError(t, err)
	assert.Equal(t, "{\"password\":\"test-password\"}", secret)
}

func TestHashiVault_GetKv1Secret_ErrorFetchingSecret(t *testing.T) {
	mockKVClient := &mockHashiKVClient{mockData: nil, mockErr: errors.New("failed to fetch secret")}
	mockVaultClient := &mockHashiVaultClientAPI{mockKVClient: mockKVClient}

	hv := &Vault{client: mockVaultClient}

	// Setup secret reference
	secretRef := secretref.New("vault", secretref.VaultTypeHashicorpVaultKv2, "secret/my-project")

	_, err := hv.GetKv1Secret(secretRef)
	require.Error(t, err)
	assert.Equal(t, "failed to fetch secret", err.Error())
}

func TestHashiVault_GetKv2Secret_Success(t *testing.T) {
	mockData := map[string]interface{}{
		"password": "test-password",
	}
	mockKVClient := &mockHashiKVClient{mockData: mockData, mockErr: nil}
	mockVaultClient := &mockHashiVaultClientAPI{mockKVClient: mockKVClient}

	hv := &Vault{client: mockVaultClient}

	// Setup secret reference
	secretRef := secretref.New("vault", secretref.VaultTypeHashicorpVaultKv2, "secret/my-project")

	secret, err := hv.GetKv2Secret(secretRef)
	require.NoError(t, err)
	assert.Equal(t, "{\"password\":\"test-password\"}", secret)
}

func TestHashiVault_GetKv2Secret_ErrorFetchingSecret(t *testing.T) {
	mockKVClient := &mockHashiKVClient{mockData: nil, mockErr: errors.New("failed to fetch secret")}
	mockVaultClient := &mockHashiVaultClientAPI{mockKVClient: mockKVClient}

	hv := &Vault{client: mockVaultClient}

	// Setup secret reference
	secretRef := secretref.New("vault", secretref.VaultTypeHashicorpVaultKv2, "secret/my-project")

	_, err := hv.GetKv2Secret(secretRef)
	require.Error(t, err)
	assert.Equal(t, "failed to fetch secret", err.Error())
}

func TestHashiVault_GetCredentials_Success(t *testing.T) {
	mockData := map[string]interface{}{
		"username": "db_username",
	}
	mockLogicalClient := &mockHashiLogicalClient{mockData: mockData, mockErr: nil}
	mockVaultClient := &mockHashiVaultClientAPI{mockLogicalClient: mockLogicalClient}

	hv := &Vault{client: mockVaultClient}

	// Setup secret reference
	secretRef := secretref.New("vault", secretref.VaultTypeHashicorpVaultKv2, "secret/my-project")

	secret, err := hv.GetCredential(secretRef)
	require.NoError(t, err)
	assert.Equal(t, "{\"username\":\"db_username\"}", secret)
}

func TestHashiVault_GetCredential_ErrorFetchingSecret(t *testing.T) {
	mockLogicalClient := &mockHashiLogicalClient{mockData: nil, mockErr: errors.New("failed to fetch secret")}
	mockVaultClient := &mockHashiVaultClientAPI{mockLogicalClient: mockLogicalClient}

	hv := &Vault{client: mockVaultClient}

	// Setup secret reference
	secretRef := secretref.New("vault", secretref.VaultTypeHashicorpVaultKv2, "secret/my-project")

	_, err := hv.GetCredential(secretRef)
	require.Error(t, err)
	assert.Equal(t, "failed to fetch secret", err.Error())
}
