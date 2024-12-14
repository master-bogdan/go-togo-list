package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func LoadConfig() *Config {
	var err = godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PORT: os.Getenv("PORT"),
	}
}
