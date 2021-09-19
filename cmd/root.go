package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "crememe",
		Short: "Remember commands.",
	}

	cmd.AddCommand(registerCmd())
	cmd.AddCommand(showCmd())
	cmd.AddCommand(removeCmd())
	cmd.AddCommand(initCmd())

	return cmd
}
