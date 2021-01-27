package main

import (
	"carsales/config"
	"carsales/database"

	"github.com/sirupsen/logrus"
)

func main() {
	config.Init()
	cfg := config.GetConfig()
	logrus.Println("Port:", cfg.Port)

	_, err := database.Connect(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

}
