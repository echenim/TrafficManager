package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWSRegion string
	ZoneID    string
}

func LoadConfig() *Config {
	err := godotenv.Load() // Load .env file if exists
	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		AWSRegion: getEnv("AWS_REGION", "us-west-2"),
		ZoneID:    getEnv("ZONE_ID", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
