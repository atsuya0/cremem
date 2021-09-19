package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/atsuya0/go-chooser"
)

const (
	pathEnv = "CREMEM_PATH"
)

type cremem struct {
	dataFile *os.File
	commands []string
	conf     Conf
}

func (c *cremem) add(command string) {
	if strings.HasPrefix(command, " ") {
		return
	}
	commandFields := strings.Fields(command)
	for _, line := range c.conf.IgnoreCommands {
		if line == commandFields[0] {
			return
		}
	}
	for _, line := range c.commands {
		if line == command {
			return
		}
	}
	c.commands = append(c.commands, command)
}

func (c *cremem) choose() (string, error) {
	errMsg := "Failed to choose commands"
	cmdChooser, err := chooser.NewChooser(c.commands)
	if err != nil {
		return "", fmt.Errorf(errMsg, err)
	}
	_, chosenCommand, err := cmdChooser.SingleRun()
	if err != nil {
		return "", fmt.Errorf(errMsg, err)
	}
	return chosenCommand, nil
}

func (c *cremem) remove() error {
	errMsg := "Failed to remove commands"
	cmdChooser, err := chooser.NewChooser(c.commands)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	chosenCommandIndexes, _, err := cmdChooser.Run()
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	for _, i := range chosenCommandIndexes {
		c.commands = append(c.commands[:i:i], c.commands[i+1:]...)
	}

	return nil
}

func (c *cremem) init() error {
	c.commands = make([]string, 0)
	return c.dataFile.Truncate(0)
}

func (c *cremem) flush() error {
	errMsg := "Failed to flash the file"
	json, err := json.MarshalIndent(Data{Commands: c.commands}, "", strings.Repeat(" ", 2))
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	if err := c.dataFile.Truncate(0); err != nil {
		return fmt.Errorf(errMsg, err)
	}
	if _, err = c.dataFile.WriteAt(json, 0); err != nil {
		return fmt.Errorf(errMsg, err)
	}

	return nil
}

func newCremem(writable bool) (cremem, error) {
	errMsg := "Failed to build cremem structure: %w"
	data, err := newData(writable)
	if err != nil {
		return cremem{}, fmt.Errorf(errMsg, err)
	}

	conf, err := newConf()
	if err != nil {
		return cremem{}, fmt.Errorf(errMsg, err)
	}

	return cremem{
		dataFile: data.file,
		commands: data.Commands,
		conf:     conf,
	}, nil
}
