package vault

import (
	"net/url"
	"strings"

	"github.com/runsecret/rsec/pkg/secretref"
)

func ConvertAwsArnToRef(arn string) string {
	// From: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	// To: rsec://123456789012/sm.aws/my-secret?region=us-west-2
	parts := strings.Split(arn, ":")
	region := parts[3]
	account := parts[4]
	secretName := parts[6]
	secretRef := secretref.New(account, secretref.VaultTypeAwsSecretsManager, secretName)
	secretRef.SetRegion(region)

	return secretRef.String()
}

func ConvertAzureArnToRef(addr string) string {
	// From: https://{vaultName}.vault.azure.net/secrets/{secretName}/{version}
	// To: rsec://myvaultname/kv.azure/mysecretname?version={version}
	parsedUrl, err := url.Parse(addr)
	if err != nil {
		return ""
	}

	// Extract the vault name and secret name from the URL
	vaultAddress := parsedUrl.Hostname()
	vaultName := strings.Split(vaultAddress, ".")[0]

	splitPath := strings.Split(parsedUrl.Path, "/")
	secretName := splitPath[2]

	secretRef := secretref.New(vaultName, secretref.VaultTypeAzureKeyVault, secretName)

	// Extract the version from the URL if present
	if len(splitPath) > 3 {
		secretVersion := splitPath[3]
		if secretVersion != "current" {
			secretRef.SetSecretVersion(secretVersion)
		}
	}

	return secretRef.String()
}
