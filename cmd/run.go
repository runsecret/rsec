/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/devenjarvis/signet/internal/core"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	// Build Command
	userCmd := exec.Command(args[0], args[1:]...)

	// Replace secrets in env vars
	env, err := core.ReplaceEnvVarSecrets(userCmd.Environ())
	if err != nil {
		fmt.Println(err)
		return
	}
	userCmd.Env = env

	// Set stdout + stderr
	userCmd.Stdout = cmd.OutOrStderr()
	userCmd.Stderr = cmd.ErrOrStderr()

	// Run command
	err = userCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
