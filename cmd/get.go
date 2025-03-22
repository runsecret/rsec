/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/devenjarvis/signet/internal/envvars"
	"github.com/spf13/cobra"
)

var OutFilePath string

// showCmd represents the show command
var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "Get the secret value from a reference",
	Long: `Get the secret value from a reference.
If the --out flag is used, the secret will be written to the specified file.`,
	Example: `  signet get my-secret
  signet get my-secret --out secret.txt`,
	Run: getFunc,
}

func getFunc(cmd *cobra.Command, args []string) {
	// Get secret based on vault type
	secretRef := args[0]
	secret, err := envvars.GetSecret(secretRef)

	// If getting the secret failed, print the error
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// If the secret is empty, print an error
	if secret == "" {
		fmt.Println("Error: Secret not found")
		return
	}

	// If the --out flag is used, write the secret to a file
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
