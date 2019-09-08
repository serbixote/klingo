package context

import (
	"fmt"
	"github.com/marco2704/klingo/internal/config"
	"github.com/spf13/cobra"
)

func newCurrentCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "current"
	cmd.Short = "Display the current context"
	cmd.Args = cobra.NoArgs
	cmd.Run = currentContext
	return cmd
}

func currentContext(cmd *cobra.Command, args []string) {
	config := config.GetKlingoConfig()
	fmt.Println(config.CurrentContext())
}
