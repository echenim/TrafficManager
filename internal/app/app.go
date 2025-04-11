package app

import (
	"dns-controller/config"
	"dns-controller/internal/dns"
	"dns-controller/internal/health"

	"github.com/sirupsen/logrus"
)

type Application struct {
	Config *config.Config
}

func NewApplication(cfg *config.Config) *Application {
	return &Application{Config: cfg}
}

func (a *Application) Run() {
	if !health.CheckEndpointHealth("https://www.stratalinks.com") {
		err := dns.UpdateDNSRecord(a.Config.ZoneID, "api.stratalinks.com", "A", "192.0.2.1")
		if err != nil {
			logrus.Error("Critical: Unable to update DNS despite endpoint being down")
		}
	}
}
