package context

import (
	"github.com/marco2704/klingo/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strings"
)

func newUseCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "use [context]"
	cmd.Short = "Set the current context, if no arg, default context is set"
	cmd.Args = cobra.MaximumNArgs(1)
	cmd.RunE = useContext
	return cmd
}

func useContext(cmd *cobra.Command, args []string) error {
	config := config.GetKlingoConfig()
	context := ""

	if len(args) == 1 {
		context = args[0]
		if strings.TrimSpace(context) == "" {
			return errors.New("invalid blank argument")
		}
	}

	return config.UseContext(context)
}
