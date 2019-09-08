package context

import (
	"github.com/spf13/cobra"
)

func NewContextCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "context"
	cmd.Short = "Manage contexts"
	cmd.Args = cobra.NoArgs
	cmd.AddCommand(
		newCreateCmd(),
		newRenameCmd(),
		newDeleteCmd(),
		newCurrentCmd(),
		newListCmd(),
		newUseCmd(),
	)

	return cmd
}
