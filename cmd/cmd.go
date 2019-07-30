package cmd

import (
	"github.com/marco2704/klingo/tui"
	"github.com/spf13/cobra"
	"log"
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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// klingo runs the UI or perform a quick connection
// if one argument is given.
func klingo(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return tui.RunTUI()
	}

	// TODO: Implement the quick connection using the args[0]
	return nil
}
