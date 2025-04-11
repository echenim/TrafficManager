package config

import (
	"os"
)

type Config struct {
	LogLevel  string
	AWSRegion string
	ZoneID    string
}

func LoadConfig() *Config {
	return &Config{
		LogLevel:  os.Getenv("LOG_LEVEL"),
		AWSRegion: os.Getenv("AWS_REGION"),
		ZoneID:    os.Getenv("ZONE_ID"),
	}
}
