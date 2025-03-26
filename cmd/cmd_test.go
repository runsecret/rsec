package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

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
