package cmd

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func isContainerRunning(containerName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "docker", "ps", "--filter", "name="+containerName, "--format", "{{.Names}}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.Contains(string(output), containerName)
}

func TestMain(m *testing.M) {
	// Set environment variables to point tests to localstack
	os.Setenv("AWS_ENDPOINT_URL", "http://localhost:4566")

	// Check if localstack container is running
	containerName := "localstack-main"
	exitCode := 1

	if isContainerRunning(containerName) {
		// Run tests if the container is running
		exitCode = m.Run()
	} else {
		// Handle the case where the container is not running
		println("Docker container", containerName, "is not running. Tests will not be executed.")
	}

	// Restore original environment variables
	os.Unsetenv("AWS_ENDPOINT_URL")

	// Exit with the test run's exit code
	os.Exit(exitCode)
}
