package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Conf struct {
	IgnoreCommands []string `json:"ignoreCommands"`
}

func getConfDir() (string, error) {
	errMsg := "Failed to get the config directory"

	if path := os.Getenv(pathEnv); path != "" {
		return path, nil
	}
	if dataHome := os.Getenv("XDG_CONFIG_HOME"); dataHome != "" {
		path := filepath.Join(dataHome, "cremem")
		if err := os.MkdirAll(path, 0700); err != nil {
			return "", fmt.Errorf(errMsg+": %w", err)
		}
		return path, nil
	}
	if homeDir, err := os.UserHomeDir(); err != nil {
		path := filepath.Join(homeDir, ".config", "cremem")
		if err := os.MkdirAll(path, 0700); err != nil {
			return "", fmt.Errorf(errMsg+": %w", err)
		}
		return path, nil
	}
	return "", errors.New(errMsg)
}

func getConfPath() (string, error) {
	path, err := getConfDir()
	if err != nil {
		return "", errors.New("Failed to get the config path")
	}
	return filepath.Join(path, "config.json"), nil
}

func newConf() (Conf, error) {
	errMsg := "Failed to build config structure: %w"

	path, err := getConfPath()
	if err != nil {
		return Conf{}, fmt.Errorf(errMsg, err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Conf{}, nil
	} else if err != nil {
		return Conf{}, fmt.Errorf(errMsg, err)
	}

	file, err := os.Open(path)
	if err != nil {
		return Conf{}, fmt.Errorf(errMsg, err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("%+v\n", fmt.Errorf(errMsg, err))
		}
	}()

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return Conf{}, fmt.Errorf(errMsg, err)
	}
	var conf Conf
	if err = json.Unmarshal(buffer.Bytes(), &conf); err != nil {
		return Conf{}, fmt.Errorf(errMsg, err)
	}

	return conf, nil
}
