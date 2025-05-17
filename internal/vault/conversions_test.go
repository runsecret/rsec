package vault

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test ConvertAwsArnToAwsRef
func TestConvertAwsArnToRef(t *testing.T) {
	testArn := "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret"
	expectedRef := "rsec://000000000000/sm.aws/MyTestSecret?region=us-east-1"
	ref := ConvertAwsArnToRef(testArn)
	assert.Equal(t, expectedRef, ref)
}

func TestConvertHashiURLToRef_KV2(t *testing.T) {
	testArn := "http://localhost:8200/v1/secret/data/my-secret"
	expectedRef := "rsec://secret/kv2.hashi/my-secret?endpoint=http%3A%2F%2Flocalhost%3A8200"
	ref, err := ConvertHashicorpVaultURLToRef(testArn)
	require.NoError(t, err)
	assert.Equal(t, expectedRef, ref)
}

func TestConvertHashiURLToRef_Creds(t *testing.T) {
	testArn := "http://localhost:8200/v1/database/creds/my-role"
	expectedRef := "rsec://database/cred.hashi/my-role?endpoint=http%3A%2F%2Flocalhost%3A8200"
	ref, err := ConvertHashicorpVaultURLToRef(testArn)
	require.NoError(t, err)
	assert.Equal(t, expectedRef, ref)
}
