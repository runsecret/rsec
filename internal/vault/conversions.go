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
	parsedURL, err := url.Parse(addr)
	if err != nil {
		return "", err
	}
	if parsedURL.Scheme != "https" {
		return "", fmt.Errorf("invalid URL scheme, azure keyvault urls must be https: %s", parsedURL.Scheme)
	}

	// Extract the vault name and secret name from the URL
	vaultAddress := parsedURL.Hostname()
	stdAzureDomain := regexp.MustCompile(`.*\.vault\.azure\.net`)
	chinaAzureDomain := regexp.MustCompile(`.*\.vault\.azure\.cn`)
	usGovAzureDomain := regexp.MustCompile(`.*\.vault\.usgovcloudapi\.net`)
	germanAzureDomain := regexp.MustCompile(`.*\.vault\.microsoftazure\.de`)

	var vaultType secretref.VaultType
	switch {
	case stdAzureDomain.MatchString(vaultAddress):
		vaultType = secretref.VaultTypeAzureKeyVault
	case chinaAzureDomain.MatchString(vaultAddress):
		vaultType = secretref.VaultTypeAzureKeyVaultChina
	case usGovAzureDomain.MatchString(vaultAddress):
		vaultType = secretref.VaultTypeAzureKeyVaultUSGov
	case germanAzureDomain.MatchString(vaultAddress):
		vaultType = secretref.VaultTypeAzureKeyVaultGermany
	default:
		return "", fmt.Errorf("invalid Azure Key Vault URL: %s", vaultAddress)
	}
	vaultName := strings.Split(vaultAddress, ".")[0]

	splitPath := strings.Split(parsedURL.Path, "/")
	secretName := splitPath[2]

	secretRef := secretref.New(vaultName, vaultType, secretName)

	// Extract the version from the URL if present
	if len(splitPath) > 3 {
		secretVersion := splitPath[3]
		if secretVersion != "current" {
			secretRef.SetSecretVersion(secretVersion)
		}
	}

	return secretRef.String(), nil
}

func ConvertHashiCorpVaultToRef(addr string) (string, error) {
	// From: https://localhost:8200/v1/{mountPath}/{data|creds}/{secretName}
	// To: rsec://{mountPath}/{kv|cred}.hashi/{secretName}
	parsedURL, err := url.Parse(addr)
	if err != nil {
		return "", err
	}

	// Extract the vault address and path
	// vaultAddress := parsedURL.Host
	vaultPath := parsedURL.Path
	// Remove v1 from the path if present
	if strings.HasPrefix(vaultPath, "/v1/") {
		vaultPath = strings.TrimPrefix(vaultPath, "/v1/")
	}

	pathSegments := strings.Split(vaultPath, "/")
	// Split the path on "data" or "creds"
	var mountPath string
	var secretName string
	var vaultType secretref.VaultType
	for i, segment := range pathSegments {
		if segment == "data" || segment == "creds" {
			switch segment {
			case "data":
				vaultType = secretref.VaultTypeHashicorpVaultKv2
			case "creds":
				vaultType = secretref.VaultTypeHashicorpVaultCred
			}
			mountPath = pathSegments[i-1]
			secretName = pathSegments[i+1]
			break
		}
	}

	// Create the secret reference
	secretRef := secretref.New(mountPath, vaultType, secretName)

	return secretRef.String(), nil
}
