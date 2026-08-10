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

func Load_json() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error finding user home directory")
		return
	}

	path := filepath.Join(home, ".gatorconfig.json")

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error reading config file:", err)
		return
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Println("error parsing JSON:", err)
		return
	}

	fmt.Println(cfg)
}
