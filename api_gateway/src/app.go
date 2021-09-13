package main

import (
	"carsales/internal/config"
	"carsales/internal/server/httpserver"
	"carsales/logger"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		log.Fatal("error in initialize application: ", err.Error())
	}
}

func run() error {
	logger := logger.New()
	defer logger.Sync()
	logger.Info("logger initialized")

	cfg := config.New()
	logger.Info("config initialized")

	server := httpserver.New()

	go func() {
		err := server.Run(cfg.HTTPServer.Domain, cfg.HTTPServer.Port)
		if err != nil {
			logger.Error("error while running http server", zap.String("err", err.Error()))
		}
	}()

	logger.Info(fmt.Sprintf("HTTP Server starting on %s:%s", cfg.HTTPServer.Domain, cfg.HTTPServer.Port))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	if err := server.Stop(); err != nil {
		logger.Error("failed stopping server", zap.String("err", err.Error()))
	}

	return nil
}
