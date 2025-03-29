package cmd

import (
	"fmt"
	"os"

	"github.com/aymanbagabas/go-osc52/v2"
	"github.com/runsecret/rsec/internal/secrets"
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
			secretRef := args[0]
			vaultClient := secrets.NewVaultClient()

			// Retrieve secret if it exists
			secret, err := vaultClient.CheckForSecret(secretRef)
			if err != nil {
				fmt.Println("X - Error retrieving secret: ", err)
			}

			// Write to clipboard using OSC 52
			if secret != "" {
				_, err = osc52.New(secret).WriteTo(os.Stderr)
				if err != nil {
					fmt.Println("X - Error writing to clipboard: ", err)
				}

				fmt.Println("✓ - Secret copied to clipboard!")
			}
			fmt.Println("X - No secret found! Double check the secret identifier provided?")
		},
	}
}

func init() {
	copyCmd := NewCopyCmd()
	rootCmd.AddCommand(copyCmd)
}
