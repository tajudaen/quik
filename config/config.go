package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	LogLevel string
}

var C Config

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	C = Config{
		Port:     os.Getenv("PORT"),
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
