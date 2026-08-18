package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	URL      string `json:"db_url"`
	Username string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error finding user home directory: %w", err)
	}
	path := filepath.Join(home, configFileName)
	return path, nil
}

func write(cfg *Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Error retreiving config file: %w", err)
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Marshaling config: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("Failed to write file: %w", err)
	}
	return nil
}

func (c *Config) SetUser(username string) error {
	c.Username = username
	err := write(c)
	if err != nil {
		return fmt.Errorf("error fetching username: %w", err)
	}
	return nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error fetching the home directory")
		return Config{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error reading config file:", err)
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Println("error parsing JSON:", err)
		return Config{}, err
	}

	return cfg, nil
}
