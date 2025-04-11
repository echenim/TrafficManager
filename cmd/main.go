package main

import (
	"dns-controller/config"
	"dns-controller/internal/app"
	"dns-controller/internal/dns"
	"dns-controller/internal/health"
	"dns-controller/pkg/logger"
	"os"
	"time"
)

func main() {
	// Initialize logging
	logger.Setup()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load configuration:", err)
		os.Exit(1)
	}

	// Create health checker with config
	healthConfig := health.NewConfig()
	healthConfig.Timeout = 5 * time.Second // Adjust timeout as needed

	// Initialize DNS manager with zone ID from config
	dnsManager := dns.NewDNSManager(cfg.ZoneID)

	// Initialize the application with dependencies
	application := app.NewApplication(cfg, dnsManager, healthConfig)

	// Run the application
	if err := application.Run(); err != nil {
		logger.Error("Application failed to run:", err)
		os.Exit(1)
	}

	// Successful execution completion
	logger.Info("Application terminated successfully")
}
