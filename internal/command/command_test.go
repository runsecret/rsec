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
