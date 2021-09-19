package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func register(cmd *cobra.Command, args []string) error {
	cremem, err := newCremem(true)
	if err != nil {
		return err
	}
	defer func() {
		if err = cremem.dataFile.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	cremem.add(args[0])
	err = cremem.flush()

	return err
}

func registerCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "register",
		Short: "Remember a command",
		Args:  cobra.MinimumNArgs(1),
		RunE:  register,
	}

	return cmd
}
