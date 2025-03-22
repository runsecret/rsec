/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/devenjarvis/signet/internal/secretref"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		secretRef := args[0]

		var convertedRef string
		switch secretref.ParseRefType(secretRef) {
		case secretref.SecretRefTypeAwsArn:
			convertedRef = secretref.ConvertAwsArnToAwsRef(secretRef)
		case secretref.SecretRefTypeAwsRef:
			convertedRef = secretref.ConvertAwsRefToAwsArn(secretRef)
		default:
			convertedRef = "Invalid secret reference"
		}

		fmt.Println(convertedRef)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
