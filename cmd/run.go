package cmd

import (
	"os/exec"

	"github.com/runsecret/rsec/internal/runsecret"
	"github.com/spf13/cobra"
)

var EnvFilePath string

func NewRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
		Short: "Run a command with secrets",
		Long: `Run a command with secrets.
If the --env flag is used, the command will be run with the environment variables loaded from the specified file.`,
		Example: `  rsec run -- echo
  rsec run --env .env -- echo`,
		Run: func(cmd *cobra.Command, args []string) {
			std := NewStd(cmd)
			// Build Command
			userCmd := exec.Command(args[0], args[1:]...)

			// Run the command with runsecret
			redactedOutput, err := runsecret.Run(userCmd, EnvFilePath)
			if err != nil {
				std.Err("Error running command: ", err)
			}

			// Output the result
			std.Out(string(redactedOutput))
		},
	}
}

func init() {
	runCmd := NewRunCmd()
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&EnvFilePath, "env", "e", "", "Env file to read env vars from")
}
