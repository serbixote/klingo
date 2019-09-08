package context

import (
	"github.com/marco2704/klingo/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strings"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "create [context]"
	cmd.Short = "Create new context"
	cmd.Args = cobra.ExactArgs(1)
	cmd.RunE = createContext
	return cmd
}

func createContext(cmd *cobra.Command, args []string) error {
	config := config.GetKlingoConfig()
	context := args[0]

	if strings.TrimSpace(context) == "" {
		return errors.New("invalid blank argument")
	}

	return config.CreateContext(context)
}
