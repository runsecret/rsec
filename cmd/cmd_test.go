package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommand(t *testing.T) {
	assert := assert.New(t)
	cmd := NewGetCmd()

	// Capture output
	b := bytes.NewBufferString("")
	cmd.SetOut(b)

	// Execute the command
	cmd.Execute()
	_, err := io.ReadAll(b)
	assert.Equal(err, nil)
}
