package main

import (
	"carsales/database/mongodb"
	"carsales/database/postgres"
	"carsales/internal/config"
	"carsales/internal/server/httpserver"
	"carsales/internal/server/routes"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	// Initialise and get config
	cfg, err := config.Init()
	if err != nil {
		return err
	}
	logrus.Println("Config: ", cfg)

	// connect and get postgres db
	pdb, err := postgres.Connect(cfg)
	if err != nil {
		return err
	}
	logrus.Println(pdb)

	// connect and get mongodb
	mdb, err := mongodb.Connect(cfg)
	if err != nil {
		return err
	}
	logrus.Println(mdb)

	// Server run
	httpserver.Run(cfg.Port, routes.Get())
	return nil
}
