package secretref

import (
	"net/url"
	"regexp"
	"strings"
)

func IsSecretRef(potentialSecret string) bool {
	refRegex := regexp.MustCompile(`(rsec:\/\/)([-a-zA-Z0-9_\+~#=:]*)\/([a-z]*\.[a-z]*)\/([-a-zA-Z0-9()@:%_\+.~#?&\/\/=]*)`)
	return refRegex.MatchString(potentialSecret)
}

type SecretReference struct {
	VaultProviderAddress string
	VaultType            VaultType
	SecretName           string
	Region               string
	SecretVersion        string
	Provider             string
}

func New(vaultProviderAddress string, vaultType VaultType, secretName string) SecretReference {
	return SecretReference{
		VaultProviderAddress: vaultProviderAddress,
		VaultType:            vaultType,
		SecretName:           secretName,
	}
}

func NewFromString(secretRef string) (SecretReference, error) {
	// Example: rsec://123456789012/sm.aws/v1/my-secret?region=us-west-2
	parsedURL, err := url.Parse(secretRef)
	if err != nil {
		return SecretReference{}, err
	}
	if parsedURL.Scheme != "rsec" {
		return SecretReference{}, err
	}
	// The host is always the vaultProviderAddress
	vaultProviderAddress := parsedURL.Hostname()

	// Get vaultType from first section of Path
	pathSegments := strings.SplitN(parsedURL.Path[1:], "/", 2)
	vaultType := pathSegments[0]

	// Extract the secret name from the path
	secretName := pathSegments[1]

	// Extract the region from the query parameters
	region := parsedURL.Query().Get("region")

	// Extract the secret version from the path if it exists
	secretVersion := parsedURL.Query().Get("version")

	// Extract the secret version from the path if it exists
	provider := parsedURL.Query().Get("provider")

	return SecretReference{
		VaultProviderAddress: vaultProviderAddress,
		VaultType:            vaultTypeFromString(vaultType),
		SecretName:           secretName,
		Region:               region,
		SecretVersion:        secretVersion,
		Provider:             provider,
	}, nil
}

func (sr *SecretReference) SetSecretVersion(version string) {
	sr.SecretVersion = version
}

func (sr *SecretReference) SetProvider(provider string) {
	sr.Provider = provider
}

func (sr *SecretReference) SetRegion(region string) {
	sr.Region = region
}

func (sr *SecretReference) String() string {
	// Example: rsec://123456789012/sm.aws/v1/my-secret?region=us-west-2
	secretRef := &url.URL{
		Scheme: "rsec",
		Host:   sr.VaultProviderAddress,
	}

	// Add the vault type
	switch sr.VaultType {
	case VaultTypeAwsSecretsManager:
		secretRef = secretRef.JoinPath("sm.aws")
	case VaultTypeAzureKeyVault:
		secretRef = secretRef.JoinPath("kv.azure")
	case VaultTypeGcpSecretsManager:
		secretRef = secretRef.JoinPath("sm.gcp")
	default:
		secretRef = secretRef.JoinPath("ERROR")
	}

	// Add secretName to the path
	secretRef = secretRef.JoinPath(sr.SecretName)

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
		return "arn:aws:secretsmanager:" + sr.Region + ":" + sr.VaultProviderAddress + ":secret:" + sr.SecretName
	case VaultTypeAzureKeyVault:
		return "https://" + sr.VaultProviderAddress + ".vault.azure.net/secrets/" + sr.SecretName
	case VaultTypeGcpSecretsManager:
		return "projects/" + sr.VaultProviderAddress + "/secrets/" + sr.SecretName
	default:
		return "Invalid vault type"
	}
}
