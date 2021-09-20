package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed cremem.zsh
var zshScriptBytes []byte

func script(cmd *cobra.Command, _ []string) error {
	fmt.Print(string(zshScriptBytes))
	return nil
}

func scriptCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "script",
		Short: "Show the zsh script to be loaded in advance.",
		Args:  cobra.NoArgs,
		RunE:  script,
	}

	return cmd
}
