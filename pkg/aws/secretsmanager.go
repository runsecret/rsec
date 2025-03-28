package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// secretName example: arn:aws:secretsmanager:us-east-2:111122223333:secret:SecretName-abcdef
func GetSecret(arn string) (string, error) {
	// Point to localstack
	awsEndpoint := "http://localhost:4566"
	// awsRegion := "us-east-1"

	awsCfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	// Create the client
	client := secretsmanager.NewFromConfig(awsCfg, func(o *secretsmanager.Options) {
		o.BaseEndpoint = aws.String(awsEndpoint)
	})

	// Get the secret
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(arn),
	}
	result, err := client.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatalf("Cannot get the secret: %s", err)
	}
	// Decrypts secret using the associated KMS CMK.
	if result.SecretString != nil {
		return *result.SecretString, nil
	}

	return "", nil
}
