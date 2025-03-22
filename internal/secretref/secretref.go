package secretref

import (
	"fmt"
	"regexp"
	"strings"
)

func GetVaultReference(secretRef string) (VaultType, string) {

	switch ParseRefType(secretRef) {
	case SecretRefTypeAwsArn:
		return VaultTypeAws, secretRef
	case SecretRefTypeAwsRef:
		return VaultTypeAws, ConvertAwsRefToAwsArn(secretRef)
	default:
		return VaultTypeUnknown, "Invalid secret reference"
	}
}

func ParseRefType(secretRef string) SecretRefType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)                      // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	awsRefRegex := regexp.MustCompile(`aws:\/\/[^\/]*\/[^\/]*\/[^\/]*`) // Ex: aws://us-west-2/123456789012/my-secret

	switch {
	case awsArnRegex.MatchString(secretRef):
		return SecretRefTypeAwsArn
	case awsRefRegex.MatchString(secretRef):
		return SecretRefTypeAwsRef
	default:
		return SecretRefTypeUnknown
	}
}

func ConvertAwsArnToAwsRef(arn string) string {
	// From: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	// To: aws://us-west-2/123456789012/my-secret
	parts := strings.Split(arn, ":")
	region := parts[3]
	account := parts[4]
	name := parts[6]
	return fmt.Sprintf("aws://%s/%s/%s", region, account, name)
}

func ConvertAwsRefToAwsArn(ref string) string {
	// From: aws://us-west-2/123456789012/my-secret
	// To: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	parts := strings.Split(ref, "/")
	region := parts[2]
	account := parts[3]
	name := parts[4]
	return fmt.Sprintf("arn:aws:secretsmanager:%s:%s:secret:%s", region, account, name)
}
