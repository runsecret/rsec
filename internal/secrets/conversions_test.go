package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test ConvertAwsArnToAwsRef
func TestConvertAwsArnToRef(t *testing.T) {
	testArn := "arn:aws:secretsmanager:us-east-1:000000000000:secret:MyTestSecret"
	expectedRef := "rsec://000000000000.sm.aws/MyTestSecret?region=us-east-1"
	ref := ConvertAwsArnToRef(testArn)
	assert.Equal(t, expectedRef, ref)
}
