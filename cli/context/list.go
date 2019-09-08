package context

import (
	"fmt"
	"github.com/marco2704/klingo/internal/config"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "list"
	cmd.Short = "List contexts marking the current one"
	cmd.Args = cobra.NoArgs
	cmd.Run = listContexts
	return cmd
}

func listContexts(cmd *cobra.Command, args []string) {
	config := config.GetKlingoConfig()
	for _, context := range config.Contexts() {
		prefix := '\u2800'
		if context == config.CurrentContext() {
			prefix = '\u2731'
		}
		fmt.Printf("%c %s\n", prefix, context)
	}
}
