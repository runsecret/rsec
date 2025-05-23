package runsecret

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/creack/pty"
	"github.com/runsecret/rsec/internal/vault"
	"github.com/runsecret/rsec/pkg/envfile"
	"github.com/runsecret/rsec/pkg/secretref"
)

func Run(userCmd *exec.Cmd, envFilePath string) ([]byte, error) {
	// Set ENV VAR secrets for the command
	envVars, secrets, err := setSecrets(userCmd, envFilePath)
	if err != nil {
		return nil, err
	}

	// Set the ENV VARs with replaced secrets on the command
	userCmd.Env = envVars

	// Run the command
	return Execute(userCmd, secrets)
}

func Execute(userCmd *exec.Cmd, secrets []string) ([]byte, error) {
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

func redactSecrets(input []byte, secretsToRedact []string) []byte {
	result := string(input)

	for _, secret := range secretsToRedact {
		// Create a replacement string of ***** to obfuscate the secret and secret length
		replacement := "*****"

		// Replace the secret with the asterisks
		result = strings.ReplaceAll(result, secret, replacement)
	}

	return []byte(result)
}

func setSecrets(cmd *exec.Cmd, envFilePath string) (envVars []string, redactList []string, err error) {
	// load ENV VARs
	envVars, err = loadEnvVars(cmd, envFilePath)
	if err != nil {
		return
	}

	// Replace secret references in ENV VARS
	vaultClient := vault.NewClient()

	for i, envVar := range envVars {
		// Split env vars
		parts := strings.SplitN(envVar, "=", 2)
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Try to get secret from env var
		var secretValue string
		if secretref.IsSecretRef(value) {
			secretValue, err = vaultClient.GetSecret(value)
			if err != nil {
				return
			}

			// Only replace secret if it is not empty
			if secretValue != "" {
				// Replace the secret in the env var
				envVars[i] = fmt.Sprintf("%s=%s", key, secretValue)
				// Add secret to list of secrets for redaction
				redactList = append(redactList, secretValue)
			}
		}
	}

	return
}

func loadEnvVars(cmd *exec.Cmd, envFilePath string) ([]string, error) {
	// Load system env vars
	cmdEnviron := cmd.Environ()

	// If --env flag used, load env vars from file
	if envFilePath != "" {
		fileEnviron, err := envfile.ReadEnvFile(envFilePath)
		if err != nil {
			return cmdEnviron, err
		}
		cmdEnviron = append(cmdEnviron, fileEnviron...)
	}

	return cmdEnviron, nil
}
