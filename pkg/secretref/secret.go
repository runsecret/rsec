package secretref

import (
	"net/url"
	"regexp"
	"strings"
)

func IsSecretReference(potentialSecret string) bool {
	refRegex := regexp.MustCompile(`(rsec:\/\/)([-a-zA-Z0-9_\+~#=]*)\.([a-z]*)\.([a-z]*)\/([-a-zA-Z0-9()@:%_\+.~#?&\/\/=]*)`)
	return refRegex.MatchString(potentialSecret)
}

type SecretReference struct {
	VaultName     string
	VaultType     VaultType
	SecretName    string
	Region        string
	SecretVersion string
}

func NewSecretReference(vaultName string, vaultType VaultType, secretName string) SecretReference {
	return SecretReference{
		VaultName:  vaultName,
		VaultType:  vaultType,
		SecretName: secretName,
	}
}

func NewSecretReferenceFromString(secretRef string) (SecretReference, error) {
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
		VaultName:     vaultName,
		VaultType:     vaultTypeFromString(vaultType),
		SecretName:    secretName,
		Region:        region,
		SecretVersion: secretVersion,
	}, nil
}

func (sr *SecretReference) SetSecretVersion(version string) {
	sr.SecretVersion = version
}

func (sr *SecretReference) SetRegion(region string) {
	sr.Region = region
}

func (sr *SecretReference) String() string {
	// Example: rsec://123456789012.sm.aws/v1/my-secret?region=us-west-2
	secretRef := url.URL{
		Scheme: "rsec",
		Host:   sr.VaultName,
	}

	// Add the vault type
	switch sr.VaultType {
	case VaultTypeAwsSecretsManager:
		secretRef.Host += ".sm.aws"
	case VaultTypeAzureKeyVault:
		secretRef.Host += ".kv.azure"
	case VaultTypeGcpSecretsManager:
		secretRef.Host += ".sm.gcp"
	default:
		secretRef.Host += ".ERROR"
	}

	secretRef.Path = "/" + sr.SecretName

	if sr.Region != "" {
		secretRef.RawQuery = "region=" + sr.Region
	}
	if sr.SecretVersion != "" {
		secretRef.Path += "?" + sr.SecretVersion
	}

	return secretRef.String()
}

func (sr *SecretReference) GetVaultAddress() string {
	// Example: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	switch sr.VaultType {
	case VaultTypeAwsSecretsManager:
		return "arn:aws:secretsmanager:" + sr.Region + ":" + sr.VaultName + ":secret:" + sr.SecretName
	case VaultTypeAzureKeyVault:
		return "https://" + sr.VaultName + ".vault.azure.net/secrets/" + sr.SecretName
	case VaultTypeGcpSecretsManager:
		return "projects/" + sr.VaultName + "/secrets/" + sr.SecretName
	default:
		return "Invalid vault type"
	}
}
