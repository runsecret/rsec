package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/runsecret/rsec/pkg/secretref"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockSecretsManagerClient is a mock implementation of the AWS Secrets Manager client interface
type MockKeyVaultClient struct {
	mock.Mock
}

// GetSecretValue mocks the GetSecretValue method
func (m *MockKeyVaultClient) GetSecret(ctx context.Context, secretName string, version string, opt *azsecrets.GetSecretOptions) (azsecrets.GetSecretResponse, error) {
	args := m.Called(ctx, secretName, version, opt)
	return args.Get(0).(azsecrets.GetSecretResponse), args.Error(1)
}

// Test the GetSecret method
func TestSecretsManager_GetSecret(t *testing.T) {
	// Create a new mock client
	mockClient := new(MockKeyVaultClient)

	// Create the service with the mock client
	service := KeyVault{mockClient}

	// Set up the mock expectation
	secretID := "test-secret"
	expectedSecret := "my-super-secret-value"
	mockClient.On("GetSecret", mock.Anything, secretID, "", mock.Anything).Return(azsecrets.GetSecretResponse{
		SecretBundle: azsecrets.SecretBundle{
			Value: &expectedSecret,
		},
	}, nil)

	// Create a secret reference
	secretRef := secretref.New("vault", secretref.VaultTypeAzureKeyVault, secretID)

	// Call the method
	secret, err := service.GetSecret(secretRef)
	require.NoError(t, err)
	assert.Equal(t, expectedSecret, secret)

	// Verify that the expected method was called
	mockClient.AssertExpectations(nil)
}
