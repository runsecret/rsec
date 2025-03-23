/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/runsecret/rsec/internal/secrets"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var refCmd = &cobra.Command{
	Use:   "ref",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "Get the secret reference and vault address from a provided secret",
	Long:  `Get the secret reference and vault address from a provided secret`,
	Example: `  rsec ref arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
  rsec ref aws://us-west-2/123456789012/my-secret`,
	Run: func(cmd *cobra.Command, args []string) {
		refOrAddr := args[0]

		var secretRef string
		var vaultAddr string
		switch secrets.GetRefType(refOrAddr) {
		case secrets.SecretRefTypeAwsArn:
			vaultAddr = refOrAddr
			secretRef = secrets.ConvertAwsArnToAwsRef(refOrAddr)
		case secrets.SecretRefTypeAwsRef:
			vaultAddr = secrets.ConvertAwsRefToAwsArn(refOrAddr)
			secretRef = refOrAddr
		default:
			secretRef = "Invalid secret address"
		}

		fmt.Println("Secret Reference: ", secretRef)
		fmt.Println("Vault Address:\t  ", vaultAddr)
	},
}

func init() {
	rootCmd.AddCommand(refCmd)
}
