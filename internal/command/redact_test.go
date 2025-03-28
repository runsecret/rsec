package command

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedactSecrets(t *testing.T) {
	// Given
	input := []byte("This is a secret: my-secret")
	secretsToRedact := []string{"my-secret"}

	// When
	result := redactSecrets(input, secretsToRedact)

	// Then
	require.Equal(t, "This is a secret: *****", string(result))
}
