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

func NewGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Short: "Get the secret value from a reference",
		Long: `Get the secret value from a reference.
If the --out flag is used, the secret will be written to the specified file.`,
		Example: `  signet get my-secret
signet get my-secret --out secret.txt`,
		RunE: func(cmd *cobra.Command, args []string) error {
			secretRef := args[0]

			// Get secret based on vault type
			secret, err := envvars.GetSecret(secretRef)
			if err != nil {
				return err
			}

			// If the --out flag is used, write the secret to a file
			if OutFilePath != "" {
				err := os.WriteFile(OutFilePath, []byte(secret), 0644)
				if err != nil {
					return err
				}
			} else {
				// Else just print it
				fmt.Println(secret)
			}
			return nil
		},
	}
}

func init() {
	getCmd := NewGetCmd()
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&OutFilePath, "out", "o", "", "File to write secret to")
}
