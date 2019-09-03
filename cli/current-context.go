package cli

import (
	"fmt"
	"github.com/marco2704/klingo/internal/config"
	"github.com/spf13/cobra"
)

func newCurrentContextCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "current-context"
	cmd.Short = "Displays the current-context"
	cmd.Args = cobra.NoArgs
	cmd.Run = currentContext
	return cmd
}

func currentContext(cmd *cobra.Command, args []string) {
	config := config.GetKlingoConfig()
	fmt.Println(config.CurrentContext())
}
