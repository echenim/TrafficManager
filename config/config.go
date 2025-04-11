package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWSRegion string
	ZoneID    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load() // Load .env file if exists
	if err != nil {
		log.Println("No .env file found")
		return nil, err.Error("")
	}

	return &Config{
		AWSRegion: getEnv("AWS_REGION", "us-west-2"),
		ZoneID:    getEnv("ZONE_ID", ""),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
