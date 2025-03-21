/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/creack/pty"
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
	env, secrets, err := core.ReplaceEnvVarSecrets(userCmd.Environ())
	if err != nil {
		fmt.Println(err)
		return
	}
	userCmd.Env = env

	// Start a pty
	ptmx, err := pty.Start(userCmd)
	if err != nil {
		fmt.Printf("Error starting pty: %v\n", err)
		return
	}
	defer ptmx.Close()

	// Read the output
	var buf bytes.Buffer
	_, err = io.Copy(&buf, ptmx)
	if err != nil {
		fmt.Printf("Error reading output: %v\n", err)
		return
	}

	// Wait for the command to finish
	if err := userCmd.Wait(); err != nil {
		fmt.Printf("Command failed: %v\n", err)
		return
	}

	// Redact secrets while preserving formatting
	rawOutput := buf.Bytes()
	redactedOutput := core.RedactSecrets(rawOutput, secrets)

	// Print the result
	fmt.Print(string(redactedOutput))
}

func init() {
	rootCmd.AddCommand(runCmd)
}
