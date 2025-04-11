package storage

import (
	"time"
)

type HealthData struct {
	URL       string
	Status    bool
	Timestamp time.Time
}

// SaveHealthCheck saves health check data
func SaveHealthCheck(data HealthData) {
	// Implementation for saving data to a database like PostgreSQL
}
