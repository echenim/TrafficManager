package health

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Config for health checks, to adjust timeouts and settings easily
type Config struct {
	Timeout time.Duration
}

// NewConfig creates a default configuration for health checks
func NewConfig() *Config {
	return &Config{
		Timeout: 5 * time.Second, // Set default timeout for health checks
	}
}

// CheckEndpoint checks the health of a given URL by making an HTTP request
func CheckEndpoint(url string, config *Config) bool {
	client := http.Client{
		Timeout: config.Timeout,
	}
	response, err := client.Get(url)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"url":   url,
			"error": err,
		}).Error("Health check failed")
		return false
	}
	defer response.Body.Close()

	// Consider the endpoint healthy if it returns HTTP 200 status
	if response.StatusCode == http.StatusOK {
		logrus.WithFields(logrus.Fields{
			"url":    url,
			"status": response.StatusCode,
		}).Info("Endpoint is healthy")
		return true
	}

	logrus.WithFields(logrus.Fields{
		"url":    url,
		"status": response.StatusCode,
	}).Warn("Endpoint returned non-200 status")
	return false
}
