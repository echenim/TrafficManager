package app

import (
	"github.com/echenim/dns-controller/config"
	"github.com/echenim/dns-controller/internal/dns"
	"github.com/echenim/dns-controller/internal/health"
	"github.com/echenim/dns-controller/pkg/logger"
)

type App struct {
	Config     *config.Config
	DNSManager *dns.DNSManager
}

func NewApp(cfg *config.Config) *App {
	dnsManager := dns.NewDNSManager(cfg.ZoneID)
	return &App{
		Config:     cfg,
		DNSManager: dnsManager,
	}
}

func (app *App) Run() {
	logger.Setup()
	if health.CheckEndpoint("https://example.com") {
		app.DNSManager.UpdateRecord("api.example.com", "A", "192.0.2.1")
	}
}
