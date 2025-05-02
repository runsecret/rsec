package cmd

import (
	"bytes"
	"io"
	"os"
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
	cmd.SetArgs([]string{"rsec://000000000000/sm.aws/test/api/keys?region=us-east-1"})

	// Execute command
	err := cmd.Execute()
	// Expect no error
	require.NoError(err)

	// Ensure output is as expected
	out, err := io.ReadAll(b)
	require.NoError(err)
	assert.Equal("✓ - Secret copied to clipboard!\n", string(out))
}

func TextCopyCommand_Azure(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	cmd := NewCopyCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Set up command arguments
	cmd.SetArgs([]string{"azure://myvaultname/mysecretname"})

	// Execute command
	err := cmd.Execute()
	// Expect no error
	require.NoError(err)

	// Ensure output is as expected
	out, err := io.ReadAll(b)
	require.NoError(err)
	assert.Equal("✓ - Secret copied to clipboard!\n", string(out))
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

func TestRunCommand(t *testing.T) {
	t.Skip("GitHub Actions doesn't like pty")
	require := require.New(t)
	assert := assert.New(t)
	cmd := NewRunCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Set up command arguments
	cmd.SetArgs([]string{"--", "echo", "password1234"})

	// Set up env vars
	os.Setenv("PASSWORD", "rsec://000000000000/sm.aws/basicPassword?region=us-east-1")

	// Execute command
	err := cmd.Execute()
	// Expect no error
	require.NoError(err)

	// Ensure output is as expected
	out, err := io.ReadAll(b)
	require.NoError(err)

	// Output is equivalent to the env var secret, and should be redacted
	assert.Equal("*****\r\n\n", string(out))

	// Clean up env vars
	os.Unsetenv("PASSWORD")
}

func TestRefCommand(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	cmd := NewRefCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Set up command arguments
	cmd.SetArgs([]string{"rsec://000000000000/sm.aws/test/api/keys?region=us-east-1"})

	// Execute command
	err := cmd.Execute()
	// Expect no error
	require.NoError(err)

	// Ensure output is as expected
	out, err := io.ReadAll(b)
	require.NoError(err)
	assert.Equal(
		"Secret Reference:  rsec://000000000000/sm.aws/test/api/keys?region=us-east-1\nVault Address:\t   arn:aws:secretsmanager:us-east-1:000000000000:secret:test/api/keys\n",
		string(out),
	)
}

func TestRefCommand_AzureArn(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	cmd := NewRefCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Set up command arguments
	cmd.SetArgs([]string{"https://myvaultname.vault.azure.net/secrets/mysecretname/"})

	// Execute command
	err := cmd.Execute()
	// Expect no error
	require.NoError(err)

	// Ensure output is as expected
	out, err := io.ReadAll(b)
	require.NoError(err)
	assert.Equal(
		"Secret Reference:  azure://myvaultname/mysecretname\nVault Address:\t   https://myvaultname.vault.azure.net/secrets/mysecretname/\n",
		string(out),
	)
}

func TestRefCommand_AzureRef(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)
	cmd := NewRefCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Set up command arguments
	cmd.SetArgs([]string{"azure://myvaultname/mysecretname"})

	// Execute command
	err := cmd.Execute()
	// Expect no error
	require.NoError(err)

	// Ensure output is as expected
	out, err := io.ReadAll(b)
	require.NoError(err)
	assert.Equal(
		"Secret Reference:  azure://myvaultname/mysecretname\nVault Address:\t   https://myvaultname.vault.azure.net/secrets/mysecretname/\n",
		string(out),
	)
}
