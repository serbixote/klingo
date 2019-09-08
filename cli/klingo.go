package cli

import (
	"fmt"
	"github.com/marco2704/klingo/cli/context"
	"github.com/marco2704/klingo/internal/tui"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:           "klingo",
		Short:         "SSH Connection Manager",
		RunE:          klingo,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootCmd.AddCommand(
		newConnectCmd(),
		context.NewContextCmd(),
	)
}

func klingo(cmd *cobra.Command, args []string) error {
	return tui.RunTUI()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
