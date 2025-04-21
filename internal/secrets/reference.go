package secrets

import (
	"net/url"
	"strings"
)

type SecretReference struct {
	vaultName     string
	vaultType     VaultType
	secretName    string
	region        string
	secretVersion string
}

func NewSecretReference(vaultName string, vaultType VaultType, secretName string) SecretReference {
	return SecretReference{
		vaultName:  vaultName,
		vaultType:  vaultType,
		secretName: secretName,
	}
}

func NewSecretReferenceFromURL(secretRef string) (SecretReference, error) {
	// Example: rsec://123456789012.sm.aws/v1/my-secret?region=us-west-2
	parsedURL, err := url.Parse(secretRef)
	if err != nil {
		return SecretReference{}, err
	}
	if parsedURL.Scheme != "rsec" {
		return SecretReference{}, err
	}
	// Extract the vault name and type from the host
	hostParts := parsedURL.Hostname()
	hostPartsSplit := strings.SplitN(hostParts, ".", 2)
	if len(hostPartsSplit) < 2 {
		return SecretReference{}, err
	}
	vaultName := hostPartsSplit[0]
	vaultType := hostPartsSplit[1]

	// Extract the secret name from the path
	secretName := parsedURL.Path[1:] // Remove leading "/"

	// Extract the region from the query parameters
	region := parsedURL.Query().Get("region")

	// Extract the secret version from the path if it exists
	secretVersion := parsedURL.Query().Get("version")

	return SecretReference{
		vaultName:     vaultName,
		vaultType:     vaultTypeFromString(vaultType),
		secretName:    secretName,
		region:        region,
		secretVersion: secretVersion,
	}, nil
}

func (sr *SecretReference) SetSecretVersion(version string) {
	sr.secretVersion = version
}

func (sr *SecretReference) SetRegion(region string) {
	sr.region = region
}

func (sr *SecretReference) String() string {
	// Example: rsec://123456789012.sm.aws/v1/my-secret?region=us-west-2
	secretRef := url.URL{
		Scheme: "rsec",
		Host:   sr.vaultName,
	}

	// Add the vault type
	switch sr.vaultType {
	case VaultTypeAwsSecretsManager:
		secretRef.Host += ".sm.aws"
	case VaultTypeAzureKeyVault:
		secretRef.Host += ".kv.azure"
	case VaultTypeGcpSecretsManager:
		secretRef.Host += ".sm.gcp"
	default:
		secretRef.Host += ".ERROR"
	}

	secretRef.Path = "/" + sr.secretName

	if sr.region != "" {
		secretRef.RawQuery = "region=" + sr.region
	}
	if sr.secretVersion != "" {
		secretRef.Path += "?" + sr.secretVersion
	}

	return secretRef.String()
}

func (sr *SecretReference) GetVaultAddress() string {
	// Example: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	switch sr.vaultType {
	case VaultTypeAwsSecretsManager:
		return "arn:aws:secretsmanager:" + sr.region + ":" + sr.vaultName + ":secret:" + sr.secretName
	case VaultTypeAzureKeyVault:
		return "https://" + sr.vaultName + ".vault.azure.net/secrets/" + sr.secretName
	case VaultTypeGcpSecretsManager:
		return "projects/" + sr.vaultName + "/secrets/" + sr.secretName
	default:
		return "Invalid vault type"
	}
}
