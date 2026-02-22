package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func createFolder() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(configPath, CONFIG_FOLDER)

	err = os.MkdirAll(fullPath, 0700)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}

func buildConfigJSON(token string) ([]byte, error) {
	config := UserConfigJSON{token}
	data, err := json.Marshal(config)

	if err != nil {
		return nil, err
	}

	return data, nil
}
