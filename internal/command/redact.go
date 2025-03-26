package command

import "strings"

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
