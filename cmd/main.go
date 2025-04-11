package main

import (
	"github.com/echenim/dns-controller/config"
	"github.com/echenim/dns-controller/internal/app"
	"github.com/echenim/dns-controller/pkg/logger"
)

func main() {
	logger.Setup()
	cfg := config.LoadConfig()
	application := app.NewApp(cfg)
	application.Run()
}
