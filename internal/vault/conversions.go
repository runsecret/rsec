package vault

import (
	"fmt"
	"net/url"
	"regexp"
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

func ConvertAzureArnToRef(addr string) (string, error) {
	// From: https://{vaultName}.vault.azure.net/secrets/{secretName}/{version}
	// To: rsec://myvaultname/kv.azure/mysecretname?version={version}&provider={provider}
	parsedUrl, err := url.Parse(addr)
	if err != nil {
		return "", err
	}
	if parsedUrl.Scheme != "https" {
		return "", fmt.Errorf("invalid URL scheme, azure keyvault urls must be https: %s", parsedUrl.Scheme)
	}

	// Extract the vault name and secret name from the URL
	vaultAddress := parsedUrl.Hostname()
	stdAzureDomain := regexp.MustCompile(`.*\.vault\.azure\.net`)
	chinaAzureDomain := regexp.MustCompile(`.*\.vault\.azure\.cn`)
	usGovAzureDomain := regexp.MustCompile(`.*\.vault\.usgovcloudapi\.net`)
	germanAzureDomain := regexp.MustCompile(`.*\.vault\.microsoftazure\.de`)

	var vaultName string
	var provider string
	if vaultAddress == "localhost" {
		vaultName = parsedUrl.Port()
	} else {
		switch {
		case stdAzureDomain.MatchString(vaultAddress):
			// Do nothing, standard vault address
		case chinaAzureDomain.MatchString(vaultAddress):
			provider = "cn"
		case usGovAzureDomain.MatchString(vaultAddress):
			provider = "usgov"
		case germanAzureDomain.MatchString(vaultAddress):
			provider = "de"
		}
		vaultName = strings.Split(vaultAddress, ".")[0]
	}

	splitPath := strings.Split(parsedUrl.Path, "/")
	secretName := splitPath[2]

	secretRef := secretref.New(vaultName, secretref.VaultTypeAzureKeyVault, secretName)

	// Add the provider to the secret reference if it exists
	if provider != "" {
		secretRef.SetProvider(provider)
	}

	// Extract the version from the URL if present
	if len(splitPath) > 3 {
		secretVersion := splitPath[3]
		if secretVersion != "current" {
			secretRef.SetSecretVersion(secretVersion)
		}
	}

	return secretRef.String(), nil
}
