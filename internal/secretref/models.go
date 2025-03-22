package secretref

type VaultType int

const (
	VaultTypeAws     VaultType = iota // AWS
	VaultTypeGcp                      // GCP - TODO: Implement
	VaultTypeAzure                    // Azure - TODO: Implement
	VaultTypeUnknown                  // Unknown
)
