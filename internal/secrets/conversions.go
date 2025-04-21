package secrets

import (
	"fmt"
	"strings"
)

func ConvertAzureArnToAzureRef(arn string) string {
	// From: https://myvaultname.vault.azure.net/secrets/mysecretname/
	// To: azure://myvaultname/mysecretname
	parts := strings.Split(arn, "/")
	vaultName := strings.Split(parts[2], ".")[0]
	secretName := parts[4]
	// TODO: Handle version if needed, for now we assume latest
	return fmt.Sprintf("azure://%s/%s", vaultName, secretName)
}

func ConvertAzureRefToAzureArn(ref string) string {
	// From: azure://myvaultname/mysecretname
	// To: https://myvaultname.vault.azure.net/secrets/mysecretname/
	parts := strings.Split(ref, "/")
	_ = parts[2]
	secretName := parts[3]
	return fmt.Sprintf("http://localhost:8080/secrets/%s/", secretName)
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
	parts := strings.SplitN(ref, "/", 5)
	region := parts[2]
	account := parts[3]
	name := parts[4]
	return fmt.Sprintf("arn:aws:secretsmanager:%s:%s:secret:%s", region, account, name)
}
