package app

import (
	"carsales/database/mongodb"
	"carsales/internal/config"

	"github.com/sirupsen/logrus"
)

// Run app
func Run(configPath string) {
	// Init config
	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Println(err)
	}

	// Connect to DB
	mongodb.Connect(cfg.Mongo.DbName, cfg.Mongo.ConnString)

	logrus.Println("Server is running on Port: ", cfg.HTTP.Port)
}
