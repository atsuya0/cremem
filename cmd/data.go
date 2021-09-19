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

type Data struct {
	file     *os.File
	Commands []string `json:"commands"`
}

func getDataDir() (string, error) {
	errMsg := "Failed to get the data directory"

	if path := os.Getenv(pathEnv); path != "" {
		return path, nil
	}
	if dataHome := os.Getenv("XDG_DATA_HOME"); dataHome != "" {
		path := filepath.Join(dataHome, "cremem")
		if err := os.MkdirAll(path, 0700); err != nil {
			return "", fmt.Errorf(errMsg+": %w", err)
		}
		return path, nil
	}
	if homeDir, err := os.UserHomeDir(); err != nil {
		path := filepath.Join(homeDir, ".local", "share", "cremem")
		if err := os.MkdirAll(path, 0700); err != nil {
			return "", fmt.Errorf(errMsg+": %w", err)
		}
		return path, nil
	}
	return "", errors.New(errMsg)
}

func getDataPath() (string, error) {
	path, err := getDataDir()
	if err != nil {
		return "", errors.New("Failed to get the data path")
	}
	return filepath.Join(path, "commands.json"), nil
}

func newData(writable bool) (Data, error) {
	errMsg := "Failed to build data structure: %w"

	path, err := getDataPath()
	if err != nil {
		return Data{}, fmt.Errorf(errMsg, err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if !writable {
			return Data{}, errors.New("Execute \"$ cremem init\"")
		}
		if file, err := os.Create(path); err != nil {
			return Data{file: file, Commands: make([]string, 0)}, fmt.Errorf(errMsg, err)
		} else {
			return Data{file: file, Commands: make([]string, 0)}, nil
		}
	} else if err != nil {
		return Data{}, fmt.Errorf(errMsg, err)
	}

	file, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return Data{}, fmt.Errorf(errMsg, err)
	}
	if !writable {
		defer func() {
			if err = file.Close(); err != nil {
				log.Fatalf("%+v\n", fmt.Errorf(errMsg, err))
			}
		}()
	}

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return Data{}, fmt.Errorf(errMsg, err)
	}
	var data Data
	if err = json.Unmarshal(buffer.Bytes(), &data); err != nil {
		return Data{}, fmt.Errorf(errMsg, err)
	}
	data.file = file

	return data, nil
}
