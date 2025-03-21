package core

type VaultType int

const (
	VaultTypeAws     VaultType = iota // AWS
	VaultTypeGcp                      // GCP
	VaultTypeAzure                    // Azure
	VaultTypeUnknown                  // Unknown
)
