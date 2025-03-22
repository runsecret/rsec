package core

import (
	"os"
	"strings"
)

func ReadEnvFile(path string) (envVars []string, err error) {
	envs, err := os.ReadFile(path)
	if err != nil {
		return
	}
	envVars = strings.Split(string(envs), "\n")

	return
}
