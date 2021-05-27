package main

import (
	"carsalesuser/internal/config"
	"carsalesuser/internal/store/postgres"
	"carsalesuser/logger"
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

	_, err := postgres.Connect(cfg)
	if err != nil {
		return err
	}
	logger.Info("Successfully connected to postgres")

	return nil
}
