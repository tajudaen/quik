package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	LogLevel      string
	MySQLUserName string
	MySQLPassword string
	MySQLHost     string
	MySQLSchema   string
	RedisURL      string
}

var C Config

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	C = Config{
		Port:          os.Getenv("PORT"),
		LogLevel:      os.Getenv("LOG_LEVEL"),
		MySQLUserName: os.Getenv("MySQL_USERNAME"),
		MySQLPassword: os.Getenv("MySQL_PASSWORD"),
		MySQLHost:     os.Getenv("MySQL_HOST"),
		MySQLSchema:   os.Getenv("MySQL_SCHEMA"),
		RedisURL:      os.Getenv("REDIS_URL"),
	}
}
