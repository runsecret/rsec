package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type std struct {
	cmd *cobra.Command
}

func NewStd(cmd *cobra.Command) std {
	return std{cmd}
}

func (s std) Out(args ...any) {
	fmt.Fprintln(s.cmd.OutOrStdout(), args...)
}

func (s std) Outf(str string, args ...any) {
	fmt.Fprintf(s.cmd.OutOrStdout(), str, args...)
}

func (s std) Err(args ...any) {
	fmt.Fprintln(s.cmd.ErrOrStderr(), args...)
}

func (s std) Errf(str string, args ...any) {
	fmt.Fprintf(s.cmd.ErrOrStderr(), str, args...)
}
