package hashi

import (
	"context"

	vault "github.com/hashicorp/vault/api"
)

// Wrapper used for HashiCorp Vault client
type vaultClientWrapper struct {
	client *vault.Client
}

func (v *vaultClientWrapper) KVv1(mountpath string) KVClient {
	return v.client.KVv1(mountpath)
}

func (v *vaultClientWrapper) KVv2(mountpath string) KVClient {
	return v.client.KVv2(mountpath)
}

func (v *vaultClientWrapper) Logical() LogicalClient {
	return v.client.Logical()
}

// Mock implementation for testing
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

type mockHashiLogicalClient struct {
	mockData map[string]any
	mockErr  error
}

func (m *mockHashiLogicalClient) Read(path string) (*vault.Secret, error) {
	if m.mockErr != nil {
		return nil, m.mockErr
	}
	return &vault.Secret{Data: m.mockData}, nil
}

type mockHashiVaultClientAPI struct {
	mockKVClient      KVClient
	mockLogicalClient LogicalClient
}

func (m *mockHashiVaultClientAPI) KVv1(mountpath string) KVClient {
	return m.mockKVClient
}

func (m *mockHashiVaultClientAPI) KVv2(mountpath string) KVClient {
	return m.mockKVClient
}

func (m *mockHashiVaultClientAPI) Logical() LogicalClient {
	return m.mockLogicalClient
}
