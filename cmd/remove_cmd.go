package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func remove(cmd *cobra.Command, _ []string) error {
	cremem, err := newCremem(true)
	if err != nil {
		return err
	}
	defer func() {
		if err = cremem.dataFile.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := cremem.remove(); err != nil {
		return err
	}
	err = cremem.flush()

	return err
}

func removeCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "remove",
		Short: "Remove commands",
		Args:  cobra.NoArgs,
		RunE:  remove,
	}

	return cmd
}
