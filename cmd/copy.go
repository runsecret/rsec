/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/runsecret/rsec/internal/envvars"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

func NewCopyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "copy",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Short: "Copy the secret value from a reference into your clipboard",
		Long:  `Copy the secret value from a reference into your clipboard`,
		Example: `  rsec copy my-secret
  rsec get my-secret --out secret.txt`,
		RunE: func(cmd *cobra.Command, args []string) error {
			secretRef := args[0]

			// Initialize clipboard
			err := clipboard.Init()
			if err != nil {
				panic(err)
			}

			// Get secret based on vault type
			secret, err := envvars.GetSecret(secretRef)
			if err != nil {
				return err
			}

			clipboard.Write(clipboard.FmtText, []byte(secret))

			fmt.Println("✓ - Secret copied to clipboard!")
			return nil
		},
	}
}

func init() {
	copyCmd := NewCopyCmd()
	rootCmd.AddCommand(copyCmd)
}
