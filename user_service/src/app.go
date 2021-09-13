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
			log.Println("not have database migrations to do")
		default:
			log.Fatal("error in initialize application: ", err.Error())
		}
	}

	log.Println("application started successfully")
}

func run() error {
	logger := logger.Get()
	defer logger.Sync()
	logger.Info("logger initialized")

	cfg := config.Get()
	logger.Info("lonfig initialized")

	pgStore, err := postgres.Connect(cfg.PgConnection)
	if err != nil {
		return err
	}
	defer pgStore.Close()
	logger.Info("successfully connected to postgres")

	err = postgres.MigrateDatabase(cfg.PgMigration, cfg.PgConnection)
	if err != nil {
		return err
	}
	logger.Info("successfully database migrations")

	return nil
}
