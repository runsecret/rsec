package utils

import "os"

func GetEnv(envVarKey string, defaultValue string) string {
	value, exists := os.LookupEnv(envVarKey)
	if !exists {
		return defaultValue
	}
	return value
}
