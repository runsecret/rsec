package aws

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSecretsManagerClient is a mock implementation of the AWS Secrets Manager client interface
type MockSecretsManagerClient struct {
	mock.Mock
}

// GetSecretValue mocks the GetSecretValue method
func (m *MockSecretsManagerClient) GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*secretsmanager.GetSecretValueOutput), args.Error(1)
}

// CreateSecret mocks the CreateSecret method
func (m *MockSecretsManagerClient) CreateSecret(ctx context.Context, params *secretsmanager.CreateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.CreateSecretOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*secretsmanager.CreateSecretOutput), args.Error(1)
}

// UpdateSecret mocks the UpdateSecret method
func (m *MockSecretsManagerClient) UpdateSecret(ctx context.Context, params *secretsmanager.UpdateSecretInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.UpdateSecretOutput, error) {
	args := m.Called(ctx, params, optFns)
	return args.Get(0).(*secretsmanager.UpdateSecretOutput), args.Error(1)
}

// Test the GetSecret method
func TestSecretsManager_GetSecret(t *testing.T) {
	// Create a new mock client
	mockClient := new(MockSecretsManagerClient)

	// Create the service with the mock client
	service := SecretsManager{"", mockClient}

	// Set up the mock expectation
	secretID := "test-secret"
	expectedSecret := "my-super-secret-value"
	mockClient.On("GetSecretValue", mock.Anything, &secretsmanager.GetSecretValueInput{
		SecretId: &secretID,
	}, mock.Anything).Return(&secretsmanager.GetSecretValueOutput{
		SecretString: &expectedSecret,
	}, nil)

	// Call the method
	secret, err := service.GetSecret(secretID)
	assert.NoError(t, err)
	assert.Equal(t, expectedSecret, secret)

	// Verify that the expected method was called
	mockClient.AssertExpectations(nil)
}

// Test the GetSecret method with an error
func TestSecretsManager_GetSecret_Error(t *testing.T) {
	// Create a new mock client
	mockClient := new(MockSecretsManagerClient)

	// Create the service with the mock client
	service := SecretsManager{"", mockClient}

	// Set up the mock expectation
	secretID := "test-secret"
	mockClient.On("GetSecretValue", mock.Anything, &secretsmanager.GetSecretValueInput{
		SecretId: &secretID,
	}, mock.Anything).Return(&secretsmanager.GetSecretValueOutput{}, assert.AnError)

	// Call the method
	secret, err := service.GetSecret(secretID)
	assert.Error(t, err)
	assert.Empty(t, secret)

	// Verify that the expected method was called
	mockClient.AssertExpectations(nil)
}
