package cli

import (
	"fmt"
	"github.com/marco2704/klingo/internal/config"
	"github.com/spf13/cobra"
)

func newGetContextsCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.Use = "get-contexts"
	cmd.Short = "Lists contexts, marking the current one"
	cmd.Args = cobra.NoArgs
	cmd.Run = getContexts
	return cmd
}

func getContexts(cmd *cobra.Command, args []string) {
	c := config.GetKlingoConfig()
	for _, context := range c.Contexts() {
		prefix := '\u2800'
		if context == c.CurrentContext() {
			prefix = '\u2731'
		}
		fmt.Printf("%c %s\n", prefix, context)
	}
}
