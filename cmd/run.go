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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
