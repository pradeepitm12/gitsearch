package config

import (
	"log"
	"os"
)

type Config struct {
	GitHubToken string
}

func Load() *Config {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN not set")
	}
	return &Config{GitHubToken: token}
}
