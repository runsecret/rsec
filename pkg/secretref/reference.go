package secretref

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/runsecret/rsec/pkg/utils"
)

func IsSecretRef(potentialSecret string) bool {
	refRegex := regexp.MustCompile(`(rsec:\/\/)([-a-zA-Z0-9_\+~#=:]*)\/([a-z0-9]*\.[a-z]*)\/([-a-zA-Z0-9()@:%_\+.~#?&\/\/=]*)`)
	return refRegex.MatchString(potentialSecret)
}

type SecretReference struct {
	VaultProviderAddress string
	VaultType            VaultType
	SecretName           string
	Region               string
	SecretVersion        string
	Endpoint             string
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
	vaultProviderAddress := parsedURL.Host

	// Get vaultType from first section of Path
	pathSegments := strings.SplitN(parsedURL.Path[1:], "/", 2)
	vaultType := pathSegments[0]

	// Extract the secret name from the path
	secretName := pathSegments[1]

	// Extract the region from the query parameters
	region := parsedURL.Query().Get("region")

	// Extract the secret version from the path if it exists
	secretVersion := parsedURL.Query().Get("version")

	// Extract the provider endpoint from the path if it exists
	endpoint := parsedURL.Query().Get("endpoint")

	return SecretReference{
		VaultProviderAddress: vaultProviderAddress,
		VaultType:            vaultTypeFromString(vaultType),
		SecretName:           secretName,
		Region:               region,
		SecretVersion:        secretVersion,
		Endpoint:             endpoint,
	}, nil
}

func (sr *SecretReference) SetSecretVersion(version string) {
	sr.SecretVersion = version
}

func (sr *SecretReference) SetRegion(region string) {
	sr.Region = region
}

func (sr *SecretReference) SetEndpoint(endpoint string) {
	sr.Endpoint = endpoint
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
	case VaultTypeAzureKeyVaultChina:
		secretRef = secretRef.JoinPath("kv.azure.cn")
	case VaultTypeAzureKeyVaultUSGov:
		secretRef = secretRef.JoinPath("kv.azure.us")
	case VaultTypeAzureKeyVaultGermany:
		secretRef = secretRef.JoinPath("kv.azure.de")
	case VaultTypeHashicorpVaultKv1:
		secretRef = secretRef.JoinPath("kv1.hashi")
	case VaultTypeHashicorpVaultKv2:
		secretRef = secretRef.JoinPath("kv2.hashi")
	case VaultTypeHashicorpVaultCred:
		secretRef = secretRef.JoinPath("cred.hashi")
	default:
		secretRef = secretRef.JoinPath("ERROR")
	}

	// Add secretName to the path
	secretRef = secretRef.JoinPath(sr.SecretName)

	//Build query
	secretRefQuery := secretRef.Query()

	if sr.Region != "" {
		secretRefQuery.Add("region", sr.Region)
	}
	if sr.Endpoint != "" {
		secretRefQuery.Add("endpoint", sr.Endpoint)
	}
	secretRef.RawQuery = secretRefQuery.Encode()

	return secretRef.String()
}

func (sr *SecretReference) GetVaultAddress() string {
	switch sr.VaultType {
	case VaultTypeAwsSecretsManager:
		return "arn:aws:secretsmanager:" + sr.Region + ":" + sr.VaultProviderAddress + ":secret:" + sr.SecretName
	case VaultTypeAzureKeyVault:
		return "https://" + sr.VaultProviderAddress + ".vault.azure.net/secrets/" + sr.SecretName
	case VaultTypeAzureKeyVaultChina:
		return "https://" + sr.VaultProviderAddress + ".vault.azure.cn/secrets/" + sr.SecretName
	case VaultTypeAzureKeyVaultUSGov:
		return "https://" + sr.VaultProviderAddress + ".vault.usgovcloudapi.net/secrets/" + sr.SecretName
	case VaultTypeAzureKeyVaultGermany:
		return "https://" + sr.VaultProviderAddress + ".vault.microsoftazure.de/secrets/" + sr.SecretName
	case VaultTypeHashicorpVaultKv1:
		vaultAddr := sr.Endpoint
		if vaultAddr == "" {
			vaultAddr = utils.GetEnv("VAULT_ADDR", "<VAULT_ADDR>")
		}
		return vaultAddr + "/v1/" + sr.VaultProviderAddress + "/" + sr.SecretName
	case VaultTypeHashicorpVaultKv2:
		vaultAddr := sr.Endpoint
		if vaultAddr == "" {
			vaultAddr = utils.GetEnv("VAULT_ADDR", "<VAULT_ADDR>")
		}
		return vaultAddr + "/v1/" + sr.VaultProviderAddress + "/data/" + sr.SecretName
	case VaultTypeHashicorpVaultCred:
		vaultAddr := sr.Endpoint
		if vaultAddr == "" {
			vaultAddr = utils.GetEnv("VAULT_ADDR", "<VAULT_ADDR>")
		}
		return vaultAddr + "/v1/" + sr.VaultProviderAddress + "/creds/" + sr.SecretName
	default:
		return "Invalid vault type"
	}
}
