package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	URL      string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.Username = username
	err := write(*c)
	if err != nil {
		return fmt.Errorf("error fetching username: %w", err)
	}
	return nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	path := filepath.Join(home, configFileName)
	return path, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error retreiving config file: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
