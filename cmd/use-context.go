package cmd

import (
	"github.com/spf13/cobra"
)

func newUseContextCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "use-context [context]"
	cmd.Short = "Sets the current context"
	cmd.Args = cobra.ExactArgs(1)
	cmd.RunE = useContext
	return cmd
}

func useContext(cmd *cobra.Command, args []string) error {
	return nil
}
