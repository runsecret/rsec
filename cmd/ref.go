package cmd

import (
	"github.com/runsecret/rsec/internal/vault"
	"github.com/runsecret/rsec/pkg/secretref"
	"github.com/spf13/cobra"
)

func NewRefCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ref",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Short: "Get the secret reference and vault address from a provided secret",
		Long:  `Get the secret reference and vault address from a provided secret`,
		Example: `  rsec ref arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
  rsec ref rsec://123456789012.sm.aws/my-secret?region=us-west-2`,
		Run: func(cmd *cobra.Command, args []string) {
			std := NewStd(cmd)
			secretID := args[0]

			var secretRef string
			var vaultAddr string
			switch vault.GetIdentifierType(secretID) {
			case vault.SecretIdentifierTypeAwsArn:
				vaultAddr = secretID
				secretRef = vault.ConvertAwsArnToRef(secretID)
			case vault.SecretIdentifierTypeRef:
				secretReference, err := secretref.NewFromString(secretID)
				if err != nil {
					std.Err("❌ - Cannot parse secret reference ", err)
					return
				}
				vaultAddr = secretReference.GetVaultAddress()
				secretRef = secretID
			default:
				secretRef = "Invalid secret identifier"
			}

			std.Out("Secret Reference: ", secretRef)
			std.Out("Vault Address:\t  ", vaultAddr)
		},
	}
}

func init() {
	refCmd := NewRefCmd()
	rootCmd.AddCommand(refCmd)
}
