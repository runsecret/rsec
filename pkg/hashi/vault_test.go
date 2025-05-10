package hashi

import (
	"context"
	"errors"
	"testing"

	vault "github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock implementations
type mockHashiKVClient struct {
	mockData map[string]interface{}
	mockErr  error
}

func (m *mockHashiKVClient) Get(ctx context.Context, path string) (*vault.KVSecret, error) {
	if m.mockErr != nil {
		return nil, m.mockErr
	}
	return &vault.KVSecret{Data: m.mockData}, nil
}

type mockHashiVaultClientAPI struct {
	mockKVClient HashiKVClient
}

func (m *mockHashiVaultClientAPI) KVv2(mountpath string) HashiKVClient {
	return m.mockKVClient
}

func TestHashiVault_GetSecret_Success(t *testing.T) {
	mockData := map[string]interface{}{
		"password": "test-password",
	}
	mockKVClient := &mockHashiKVClient{mockData: mockData, mockErr: nil}
	mockVaultClient := &mockHashiVaultClientAPI{mockKVClient: mockKVClient}

	hv := &HashiVault{client: mockVaultClient}

	secret, err := hv.GetSecret("http://localhost:8200", "secret/my-project")
	require.NoError(t, err)
	assert.Equal(t, "test-password", secret)
}

func TestHashiVault_GetSecret_ErrorFetchingSecret(t *testing.T) {
	mockKVClient := &mockHashiKVClient{mockData: nil, mockErr: errors.New("failed to fetch secret")}
	mockVaultClient := &mockHashiVaultClientAPI{mockKVClient: mockKVClient}

	hv := &HashiVault{client: mockVaultClient}

	_, err := hv.GetSecret("http://localhost:8200", "secret/my-project")
	require.Error(t, err)
	assert.Equal(t, "failed to fetch secret", err.Error())
}

func TestHashiVault_GetSecret_InvalidDataType(t *testing.T) {
	mockData := map[string]interface{}{
		"password": 12345, // Invalid type
	}
	mockKVClient := &mockHashiKVClient{mockData: mockData, mockErr: nil}
	mockVaultClient := &mockHashiVaultClientAPI{mockKVClient: mockKVClient}

	hv := &HashiVault{client: mockVaultClient}
	_, err := hv.GetSecret("http://localhost:8200", "secret/my-project")
	require.Error(t, err)
	assert.Equal(t, "value type assertion failed: int 12345", err.Error())
}
