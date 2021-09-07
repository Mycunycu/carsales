package main

import (
	"carsales/internal/config"
	"carsales/internal/server/httpserver"
	"carsales/logger"
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("Error in initialize application: ", err.Error())
	}
}

func run() error {
	logger := logger.New()
	defer logger.Sync()
	logger.Info("Logger initialized")

	cfg := config.New()
	logger.Info("Config initialized")

	logger.Info(fmt.Sprintf("HTTP Server starting on %s:%s", cfg.HTTPServer.Domain, cfg.HTTPServer.Port))
	err := httpserver.Run(cfg.HTTPServer.Domain, cfg.HTTPServer.Port)
	if err != nil {
		return err
	}

	return nil
}
