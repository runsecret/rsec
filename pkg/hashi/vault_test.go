package hashi

import (
	"context"
	"errors"
	"testing"

	vault "github.com/hashicorp/vault/api"
	"github.com/runsecret/rsec/pkg/secretref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock implementations
type mockHashiKVClient struct {
	mockData map[string]any
	mockErr  error
}

func (m *mockHashiKVClient) Get(ctx context.Context, path string) (*vault.KVSecret, error) {
	if m.mockErr != nil {
		return nil, m.mockErr
	}
	return &vault.KVSecret{Data: m.mockData}, nil
}

type mockHashiVaultClientAPI struct {
	mockKVClient KVClient
}

func (m *mockHashiVaultClientAPI) KVv1(mountpath string) KVClient {
	return m.mockKVClient
}

func (m *mockHashiVaultClientAPI) KVv2(mountpath string) KVClient {
	return m.mockKVClient
}

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
