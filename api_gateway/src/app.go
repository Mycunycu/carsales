package main

import (
	"carsales/internal/config"
	"carsales/internal/server/httpserver"
	"carsales/internal/server/routes"
	"carsales/logger"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("Error in initialize application: ", err)
	}
}

func run() error {
	logger := logger.Get()
	defer logger.Sync()
	logger.Info("Logger initialized")

	cfg := config.Get()
	logger.Info("Config initialized")

	httpserver.Run(cfg.Server.Port, routes.Get())
	return nil
}
