package app

import (
	"carsales/database/mongodb"
	"carsales/internal/config"
	"carsales/internal/server/httpserver"
	"carsales/internal/server/routes"

	"github.com/sirupsen/logrus"
)

// Run app
func Run(configPath string) {
	// Init config
	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Fatal(err)
	}

	// Connect to DB
	store := new(mongodb.Store)
	store.Connect(cfg.Mongo.DbName, cfg.Mongo.ConnString)

	// Server run
	srv := new(httpserver.Server)
	router := new(routes.Routes)
	srv.Run(cfg.HTTP.Port, router.Init())

	logrus.Println("Server is running on Port: ", cfg.HTTP.Port)
}
