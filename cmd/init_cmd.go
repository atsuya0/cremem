package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func initialize(cmd *cobra.Command, _ []string) error {
	cremem, err := newCremem(true)
	if err != nil {
		return err
	}
	defer func() {
		if err = cremem.dataFile.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	return cremem.init()
}

func initCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "initialize the data",
		Args:  cobra.NoArgs,
		RunE:  initialize,
	}

	return cmd
}
