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
		RunE: func(cmd *cobra.Command, args []string) error {
			secretRef := args[0]

			// Get secret based on vault type
			secret, err := secrets.GetSecret(secretRef)
			if err != nil {
				return err
			}

			// Write to clipboard using OSC 52
			_, err = osc52.New(secret).WriteTo(os.Stderr)
			if err != nil {
				return err
			}

			fmt.Println("✓ - Secret copied to clipboard!")
			return nil
		},
	}
}

func init() {
	copyCmd := NewCopyCmd()
	rootCmd.AddCommand(copyCmd)
}
