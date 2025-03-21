/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/devenjarvis/signet/internal/aws"
	"github.com/devenjarvis/signet/internal/core"
	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: read,
}

func read(cmd *cobra.Command, args []string) {
	// Identify type of secret vault
	secretRef := args[0]
	vaultType := core.ParseVualtType(secretRef)

	var secret string
	var err error

	// Get secret based on vault type
	switch vaultType {
	case core.VaultTypeAws:
		secret, err = aws.GetSecret(args[0])
	default:
		fmt.Println("Error: unimplemented vault type")
	}

	// If getting the secret failed, print the error
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	// Else print the secret
	fmt.Println(secret)
}

func init() {
	rootCmd.AddCommand(readCmd)
}
