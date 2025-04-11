package app

import (
	"github.com/echenim/dns-controller/config"
	"github.com/echenim/dns-controller/internal/dns"
	"github.com/echenim/dns-controller/internal/health"

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
