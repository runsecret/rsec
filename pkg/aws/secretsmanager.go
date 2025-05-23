package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// Define the interface to match the AWS Secrets Manager client methods
type SecretsManagerAPI interface {
	GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

type SecretsManager struct {
	awsEndpoint string
	client      SecretsManagerAPI
}

func NewSecretsManager() *SecretsManager {
	return &SecretsManager{}
}

func (s *SecretsManager) getClient() SecretsManagerAPI {
	// Create client if not already created
	if s.client == nil {
		awsCfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("Cannot load the AWS configs: %s", err)
		}

		// Create the client
		client := secretsmanager.NewFromConfig(awsCfg, func(o *secretsmanager.Options) {})
		s.client = client
	}

	// Return the client
	return s.client
}

func (s SecretsManager) GetSecret(arn string) (string, error) {
	// Create input for the GetSecretValue API call
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(arn),
	}

	// Call the GetSecretValue API
	result, err := s.getClient().GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", err
	}

	// Decrypt and return secret string or binary
	if result.SecretString != nil {
		return *result.SecretString, nil
	}
	// If the secret is binary, return an empty string
	if result.SecretBinary != nil {
		return "", nil
	}

	return "", nil
}
