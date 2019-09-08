package context

import (
	"github.com/marco2704/klingo/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strings"
)

func newRenameCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "rename [old-context] [new-context]"
	cmd.Short = "Rename context"
	cmd.Args = cobra.ExactArgs(2)
	cmd.RunE = renameContext
	return cmd
}

func renameContext(cmd *cobra.Command, args []string) error {
	config := config.GetKlingoConfig()
	oldContext := args[0]
	newContext := args[1]

	if strings.TrimSpace(oldContext) == "" || strings.TrimSpace(newContext) == "" {
		return errors.New("invalid blank argument")
	}

	return config.RenameContext(oldContext, newContext)
}
