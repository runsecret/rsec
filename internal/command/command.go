package command

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/creack/pty"
)

func Run(userCmd *exec.Cmd, secrets []string) ([]byte, error) {
	// Start a pty
	ptmx, err := pty.Start(userCmd)
	if err != nil {
		return nil, fmt.Errorf("error starting pty: %v", err)
	}
	defer ptmx.Close()

	// Read the output
	var buf bytes.Buffer
	_, err = io.Copy(&buf, ptmx)
	if err != nil {
		return nil, fmt.Errorf("error reading output: %v", err)
	}

	// Wait for the command to finish
	if err := userCmd.Wait(); err != nil {
		return nil, fmt.Errorf("command failed: %v", err)
	}

	// Redact secrets while preserving formatting
	rawOutput := buf.Bytes()
	return redactSecrets(rawOutput, secrets), nil
}
