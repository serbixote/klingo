package context

import (
	"github.com/marco2704/klingo/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strings"
)

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "delete [context]"
	cmd.Short = "Delete context"
	cmd.Args = cobra.ExactArgs(1)
	cmd.RunE = deleteContext
	return cmd
}

func deleteContext(cmd *cobra.Command, args []string) error {
	config := config.GetKlingoConfig()
	context := args[0]

	if strings.TrimSpace(context) == "" {
		return errors.New("invalid blank argument")
	}

	return config.DeleteContext(context)
}
