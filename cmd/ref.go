package cmd

import (
	"github.com/runsecret/rsec/internal/secrets"
	"github.com/spf13/cobra"
)

func NewRefCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ref",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Short: "Get the secret reference and vault address from a provided secret",
		Long:  `Get the secret reference and vault address from a provided secret`,
		Example: `  rsec ref arn:aws:secretsmanager:us-west-2:123456789012:secret:my-secret
  rsec ref aws://us-west-2/123456789012/my-secret`,
		Run: func(cmd *cobra.Command, args []string) {
			std := NewStd(cmd)
			refOrAddr := args[0]

			var secretRef string
			var vaultAddr string
			switch secrets.GetIdentifierType(refOrAddr) {
			case secrets.SecretIdentifierTypeAwsArn:
				vaultAddr = refOrAddr
				secretRef = secrets.ConvertAwsArnToAwsRef(refOrAddr)
			case secrets.SecretIdentifierTypeAwsRef:
				vaultAddr = secrets.ConvertAwsRefToAwsArn(refOrAddr)
				secretRef = refOrAddr
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
