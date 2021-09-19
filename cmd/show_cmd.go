package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func show(cmd *cobra.Command, _ []string) error {
	cremem, err := newCremem(false)
	if err != nil {
		return err
	}

	command, err := cremem.choose()
	if err != nil {
		return err
	} else if command != "" {
		fmt.Println(command)
	}
	return nil
}

func showCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "show",
		Short: "Show a command",
		Args:  cobra.NoArgs,
		RunE:  show,
	}

	return cmd
}
