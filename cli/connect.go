package cli

import (
	"github.com/spf13/cobra"
)

func newConnectCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "connect [connection]"
	cmd.Short = "Connect to a remote machine"
	cmd.Args = cobra.ExactArgs(1)
	cmd.RunE = connect
	return cmd
}

func connect(cmd *cobra.Command, args []string) error {
	return nil
}
