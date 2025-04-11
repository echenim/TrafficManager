package app

import (
	"time"

	"github.com/echenim/dns-controller/internal/analytics"
	"github.com/echenim/dns-controller/internal/dns"
	"github.com/echenim/dns-controller/internal/health"
	"github.com/echenim/dns-controller/internal/storage"
)

type App struct {
	DNSManager *dns.DNSManager
}

func (app *App) Run() {
	urls := []string{"https://example.com", "https://example2.com"}
	for _, url := range urls {
		if health.CheckEndpoint(url) {
			if !analytics.PredictHealth(url) {
				app.DNSManager.UpdateRecord("ZN00234", "api.example.com", "A", "192.0.2.1")
			}
		} else {
			storage.SaveHealthCheck(storage.HealthData{URL: url, Status: false, Timestamp: time.Now()})
		}
	}
}
