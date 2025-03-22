/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/devenjarvis/signet/internal/aws"
	"github.com/devenjarvis/signet/internal/envvars"
	"github.com/spf13/cobra"
)

var OutFilePath string

// showCmd represents the show command
var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: get,
}

func get(cmd *cobra.Command, args []string) {
	// Identify type of secret vault
	secretRef := args[0]
	vaultType := envvars.ParseVaultType(secretRef)

	var secret string
	var err error

	// Get secret based on vault type
	switch vaultType {
	case envvars.VaultTypeAws:
		secret, err = aws.GetSecret(args[0])
	default:
		fmt.Println("Error: unimplemented vault type")
	}

	// If getting the secret failed, print the error
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	// If the --out flag is used write the secret to a file
	if OutFilePath != "" {
		err := os.WriteFile(OutFilePath, []byte(secret), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	} else {
		// Else just print it
		fmt.Println(secret)
	}
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&OutFilePath, "out", "o", "", "File to write secret to")
}
