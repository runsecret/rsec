package cmd

import (
	"fmt"
	"os"

	"github.com/aymanbagabas/go-osc52/v2"
	"github.com/runsecret/rsec/internal/vault"
	"github.com/runsecret/rsec/pkg/secretref"
	"github.com/spf13/cobra"
)

func NewCopyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "copy",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Short: "Copy the secret value from a reference into your clipboard",
		Long:  `Copy the secret value from a reference into your clipboard`,
		Example: `  rsec copy my-secret
  rsec get my-secret --out secret.txt`,
		Run: func(cmd *cobra.Command, args []string) {
			std := NewStd(cmd)
			secretRef := args[0]
			if !secretref.IsSecretReference(secretRef) {
				std.Err("❌ - Invalid secret reference provided!")
				return
			}

			vaultClient := vault.NewClient()

			// Retrieve secret if it exists
			secret, err := vaultClient.CheckForSecret(secretRef)
			// Error handling
			if err != nil {
				std.Err("❌ - Error retrieving secret: ", err)
				return
			}
			if secret == "" {
				std.Err("X - No secret found! Double check the secret identifier provided?")
				return
			}

			// Write to clipboard using OSC 52
			_, err = osc52.New(secret).WriteTo(os.Stderr)
			if err != nil {
				fmt.Println("❌ - Error writing to clipboard: ", err)
			}

			std.Out("✓ - Secret copied to clipboard!")
		},
	}
}

func init() {
	copyCmd := NewCopyCmd()
	rootCmd.AddCommand(copyCmd)
}
