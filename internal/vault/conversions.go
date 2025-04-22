package vault

import (
	"strings"

	"github.com/runsecret/rsec/pkg/secret"
)

func ConvertAwsArnToRef(arn string) string {
	// From: arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
	// To: rsec://123456789012.sm.aws/v1/my-secret?region=us-west-2
	parts := strings.Split(arn, ":")
	region := parts[3]
	account := parts[4]
	secretName := parts[6]
	secretRef := secret.NewSecretReference(account, secret.VaultTypeAwsSecretsManager, secretName)
	secretRef.SetRegion(region)

	return secretRef.String()
}
