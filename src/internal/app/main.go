package main

import (
	"carsales/database/mongodb"
	"carsales/database/postgres"
	"carsales/internal/config"
	"carsales/internal/server/httpserver"
	"carsales/internal/server/routes"
	"carsales/pkg/logger"
	"log"

	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("Error in initialize application: ", err)
	}
}

func run() error {
	// Initialise and get config
	cfg := config.Init()
	logger := logger.Get()
	defer logger.Sync()

	logger.Info("Config initialized", zap.Any("Config", cfg))

	// connect and get postgres db
	pgStore, err := postgres.Connect(cfg)
	if err != nil {
		return err
	}
	logger.Info("Postgres connected", zap.Any("PostgesStore", pgStore))

	// connect and get mongodb
	mongoStore, err := mongodb.Connect(cfg)
	if err != nil {
		return err
	}
	logger.Info("Postgres connected", zap.Any("MongoStore", mongoStore))

	// Server run
	httpserver.Run(cfg.Port, routes.Get())
	return nil
}
