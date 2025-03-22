package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCommand(t *testing.T) {
	require := require.New(t)
	cmd := NewGetCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Execute the command
	err := cmd.Execute()
	require.Error(err)
	_, err = io.ReadAll(b)
	require.NoError(err)
}
