package main

import (
	"log"
	"os"

	"github.com/echenim/dns-controller/config"
	"github.com/echenim/dns-controller/internal/analytics"
	"github.com/echenim/dns-controller/internal/app"
	"github.com/echenim/dns-controller/internal/dns"
	"github.com/echenim/dns-controller/internal/health"
	"github.com/echenim/dns-controller/internal/storage"
	"github.com/echenim/dns-controller/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// Set up the logging configuration
	logger.Setup()

	// Load the application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the DNS Manager
	dnsManager := dns.NewDNSManager(cfg.ZoneID)

	// Initialize the App with all components
	application := app.NewApp(cfg, dnsManager, health.CheckEndpoint, analytics.PredictHealth, storage.SaveHealthCheck)

	// Run the application
	if err := application.Run(); err != nil {
		logger.Error("Application failed to run:", err)
		os.Exit(1)
	}

	// If the application exits without errors
	logger.Info("Application terminated successfully")
}
