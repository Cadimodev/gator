package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(home, ConfigFilePath, ConfigFileName)
	return fullPath, nil
}

func Read() (Config, error) {

	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	configData := Config{}
	err = json.Unmarshal(data, &configData)
	if err != nil {
		return Config{}, err
	}

	return configData, nil
}

func (cfg *Config) SetUser(user string) error {

	cfg.CurrentUserName = user

	err := write(cfg)
	if err != nil {
		return err
	}

	return nil
}

func write(cfg *Config) error {

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
