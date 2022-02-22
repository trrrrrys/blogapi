package config

import "os"

type Config struct {
	ProjectID string
	Port      string
}

func NewConfig() *Config {
	return &Config{
		ProjectID: os.Getenv("PROJECT_ID"),
		// Port:      os.Getenv("PORT"),
		Port: "8080",
	}
}
