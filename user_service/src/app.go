package main

import (
	"carsales_user/internal/config"
	"carsales_user/logger"
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

	logger.Info(cfg.RPCServer.Port)

	return nil
}
