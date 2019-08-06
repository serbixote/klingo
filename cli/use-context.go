package cli

import (
	"errors"
	"github.com/marco2704/klingo/internal/config"
	"github.com/spf13/cobra"
	"strings"
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
	config, err := config.GetKlingoConfig()
	if err != nil {
		return err
	}

	context := args[0]

	if strings.TrimSpace(context) == "" {
		return errors.New("Invalid blank argument")
	}

	return config.UseContext(context)
}
