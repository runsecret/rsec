/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/runsecret/rsec/internal/tui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rsync",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Short: "The safe and easy way to use secrets on your local machine",
	Long: `rsec is a CLI tool that helps you use secrets safely on your local machine.
Set environment variables with secret references and rsec will replace them with the actual secret values at run time.
rsec also redacts secrets from the output of commands, minimizing the possible exposure of sensitive information.
`,
	Run: runFunc,
}

func RenderUsage(cmd *cobra.Command) error {
	fmt.Println(tui.RenderMkDown(cmd.Use))
	return nil
}

func RenderHelp(cmd *cobra.Command, smthing []string) {
	fmt.Println(tui.RenderMkDown(cmd.Long))
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd.SetUsageFunc(RenderUsage)
	// rootCmd.SetHelpFunc(RenderHelp)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
