package config

import "io"

type Config struct {
	URL      string `json:"db_url"`
	Username string `json:"current_user_name"`
}
