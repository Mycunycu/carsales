package logger

import (
	"carsales/internal/config"

	"go.uber.org/zap"
)

func Get() *zap.Logger {
	var logger *zap.Logger
	cfg := config.GetConfig()

	if cfg.Env == "prod" {
		logger, _ = zap.NewProduction()
		logger.Info("Will used PROD preset logger")
	}

	logger, _ = zap.NewDevelopment()
	logger.Info("Will used DEV preset logger")

	return logger
}
