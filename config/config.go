package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TgToken string
	TgHost  string
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	return &Config{
		TgToken: os.Getenv("TG_TOKEN"),
		TgHost:  os.Getenv("TG_HOST"),
	}, nil
}
