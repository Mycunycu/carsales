package main

import (
	"carsales/config"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Init()
	cfg := config.GetConfig()
	logrus.Printf("%s\n", cfg.Port)
}
