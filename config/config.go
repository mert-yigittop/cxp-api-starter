package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var config Config
	config.DBHost = getEnv("DB_HOST", "localhost")
	config.DBUser = getEnv("DB_USER", "postgres")
	config.DBPassword = getEnv("DB_PASSWORD", "password")
	config.DBName = getEnv("DB_NAME", "todo")
	config.DBPort = getEnv("DB_PORT", "5432")

	AppConfig = &config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
