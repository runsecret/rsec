/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "signet",
	Short: "The safe and easy way to use secrets on your local machine",
	Long: `Signet is a CLI tool that helps you use secrets safely on your local machine. 
Set environment variables with secret references and Signet will replace them with the actual secret values at run time.
Signet also redacts secrets from the output of commands, minimizing the possible exposure of sensitive information.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
