package main

import (
	"carsalesuser/internal/config"
	"carsalesuser/internal/store/postgres"
	"carsalesuser/logger"
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	if err := run(); err != nil {
		switch err {
		case migrate.ErrNoChange:
			log.Println("Not have database migrations to do")
		default:
			log.Fatal("Error in initialize application: ", err.Error())
		}
	}

	log.Println("Application started successfully")
}

func run() error {
	logger := logger.Get()
	defer logger.Sync()
	logger.Info("Logger initialized")

	cfg := config.Get()
	logger.Info("Config initialized")

	pgStore, err := postgres.Connect(cfg.PgConnection)
	if err != nil {
		return err
	}
	defer pgStore.Close()
	logger.Info("Successfully connected to postgres")

	err = postgres.MigrateDatabase(cfg.PgMigration, cfg.PgConnection)
	if err != nil {
		return err
	}
	logger.Info("Successfully database migrations")

	return nil
}
