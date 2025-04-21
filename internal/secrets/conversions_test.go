package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test ConvertAwsArnToAwsRef
func TestConvertAwsArnToAwsRef(t *testing.T) {
	testArn := "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret"
	expectedRef := "aws://us-east-1/000000000000/MyTestSecret"
	ref := ConvertAwsArnToAwsRef(testArn)
	assert.Equal(t, expectedRef, ref)
}

// Test ConvertAwsRefToAwsArn
func TestConvertAwsRefToAwsArn(t *testing.T) {
	testRef := "aws://us-east-1/000000000000/MyTestSecret"
	expectedArn := "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret"
	arn := ConvertAwsRefToAwsArn(testRef)
	assert.Equal(t, expectedArn, arn)
}

// Test ConvertAzureArnToAzureRef
func TestConvertAzureArnToAzureRef(t *testing.T) {
	testArn := "https://myvaultname.vault.azure.net/secrets/mysecretname/"
	expectedRef := "azure://myvaultname/mysecretname"
	ref := ConvertAzureArnToAzureRef(testArn)
	assert.Equal(t, expectedRef, ref)
}

// Test ConvertAzureRefToAzureArn
func TestConvertAzureRefToAzureArn(t *testing.T) {
	testRef := "azure://myvaultname/mysecretname"
	expectedArn := "https://myvaultname.vault.azure.net/secrets/mysecretname/"
	arn := ConvertAzureRefToAzureArn(testRef)
	assert.Equal(t, expectedArn, arn)
}
