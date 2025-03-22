/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/devenjarvis/signet/internal/command"
	"github.com/devenjarvis/signet/internal/envfile"
	"github.com/devenjarvis/signet/internal/envvars"
	"github.com/spf13/cobra"
)

var EnvFilePath string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command with secrets",
	Long: `Run a command with secrets.
If the --env flag is used, the command will be run with the environment variables loaded from the specified file.`,
	Example: `signet run -- echo
signet run --env .env -- echo`,
	Run: runFunc,
}

func loadEnvVars(userCmd *exec.Cmd) ([]string, error) {
	// Load system env vars
	cmdEnviron := userCmd.Environ()

	// If --env flag used, load env vars from file
	if EnvFilePath != "" {
		fileEnviron, err := envfile.Read(EnvFilePath)
		if err != nil {
			return cmdEnviron, err
		}
		cmdEnviron = append(cmdEnviron, fileEnviron...)
	}

	return cmdEnviron, nil
}

func runFunc(cmd *cobra.Command, args []string) {
	// Build Command
	userCmd := exec.Command(args[0], args[1:]...)

	// Load env vars
	cmdEnviron, err := loadEnvVars(userCmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Replace secrets in env vars
	env, secrets, err := envvars.SetSecrets(cmdEnviron)
	if err != nil {
		fmt.Println(err)
		return
	}
	userCmd.Env = env

	// Run Command
	redactedOutput, err := command.Run(userCmd, secrets)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the result
	fmt.Print(string(redactedOutput))
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&EnvFilePath, "env", "e", "", "Env file to read env vars from")
}
