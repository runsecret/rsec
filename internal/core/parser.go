package core

import (
	"regexp"
)

func ParseVualtType(secretRef string) VaultType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)

	switch {
	case awsArnRegex.MatchString(secretRef):
		return VaultTypeAws
	default:
		return VaultTypeUnknown
	}

}
