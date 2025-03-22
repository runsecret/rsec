package secretref

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseVaultType(secretRef string) VaultType {
	awsArnRegex := regexp.MustCompile(`arn:aws.*`)                      // Ex: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	awsRefRegex := regexp.MustCompile(`aws:\/\/[^\/]*\/[^\/]*\/[^\/]*`) // Ex: aws://us-west-2/123456789012/my-secret

	switch {
	case awsArnRegex.MatchString(secretRef), awsRefRegex.MatchString(secretRef):
		return VaultTypeAws
	default:
		return VaultTypeUnknown
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
