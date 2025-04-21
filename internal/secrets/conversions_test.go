package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test ConvertAwsArnToAwsRef
func TestConvertAwsArnToRef(t *testing.T) {
	testArn := "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret"
	expectedRef := "rsec://000000000000.sm.aws/MyTestSecret?region=us-east-1"
	ref := ConvertAwsArnToRef(testArn)
	assert.Equal(t, expectedRef, ref)
}

// Test ConvertAwsRefToAwsArn
func TestConvertRefToAwsArn(t *testing.T) {
	testRef := "rsec://000000000000.sm.aws/MyTestSecret?region=us-east-1"
	expectedArn := "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret"
	arn, err := ConvertRefToAwsArn(testRef)
	require.NoError(t, err)
	assert.Equal(t, expectedArn, arn)
}
