package command

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	// Given (ls formats output differently if it detects a TTY)
	cmd := exec.Command("ls")
	secrets := []string{}

	// When
	output, err := Run(cmd, secrets)

	// Then
	require.NoError(t, err)
	require.Equal(t, "command.go\tcommand_test.go\tredact.go\tredact_test.go\r\n", string(output))
}

func TestRedactSecrets(t *testing.T) {
	// Given
	input := []byte("This is a secret: my-secret")
	secretsToRedact := []string{"my-secret"}

	// When
	result := redactSecrets(input, secretsToRedact)

	// Then
	require.Equal(t, "This is a secret: *****", string(result))
}
