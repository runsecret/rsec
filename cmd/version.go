package cmd

import (
	"github.com/spf13/cobra"
)

var version string

func SetVersion(v string) {
	version = v
}

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Returns the currently installed version of rsec",
		Long:    `Returns the currently installed version of rsec`,
		Example: `  rsec version`,
		Run: func(cmd *cobra.Command, args []string) {
			std := NewStd(cmd)
			std.Out(version)
		},
	}
}

func init() {
	versionCmd := NewVersionCmd()
	rootCmd.AddCommand(versionCmd)
}
