package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SetToken(token string) error {
	path, err := createFolder()
	if err != nil {
		return err
	}

	bytes, err := buildConfigJSON(token)

	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(path, ConfigFileName), bytes, 0700)

	if err != nil {
		return err
	}

	return nil
}

func GetConfig() (UserConfigJSON, error) {
	configPath, err := os.UserConfigDir()
	userConfig := UserConfigJSON{}

	if err != nil {
		return userConfig, err
	}

	configFile, err := os.ReadFile(filepath.Join(configPath, ConfigFolder, ConfigFileName))

	if err != nil {
		return userConfig, err
	}

	err = json.Unmarshal(configFile, &userConfig)

	if err != nil {
		return userConfig, err
	}

	return userConfig, nil
}
