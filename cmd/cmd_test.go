package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCopyCommand(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	cmd := NewCopyCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Set up command arguments
	cmd.SetArgs([]string{"aws://us-east-1/000000000000/test/api/keys"})

	// Execute command
	err := cmd.Execute()
	// Expect no error
	require.NoError(err)

	// Ensure output is as expected
	out, err := io.ReadAll(b)
	require.NoError(err)
	assert.Equal(string(out), "✓ - Secret copied to clipboard!\n")
}

func TestCopyCommand_MissingArgument(t *testing.T) {
	require := require.New(t)
	cmd := NewCopyCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Execute command with no arguments
	err := cmd.Execute()
	// Expect an error
	require.Error(err)

	// Ensure output is as expected
	_, err = io.ReadAll(b)
	require.NoError(err)
}
