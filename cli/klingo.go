package cli

import (
	"fmt"
	"github.com/marco2704/klingo/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:           "klingo",
		Short:         "SSH Connection Manager",
		Args:          cobra.MaximumNArgs(1),
		RunE:          klingo,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootCmd.AddCommand(
		newUseContextCmd(),
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func klingo(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return tui.RunTUI()
	}

	// TODO: Implement the quick connection using the args[0]
	return nil
}
