package logger

import (
	"carsales/internal/config"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func Get() Logger {
	var zapLogger *zap.Logger
	cfg := config.Get()

	if cfg.Env == config.PROD_ENV {
		zapLogger, _ = zap.NewProduction()
		zapLogger.Info("Will used PROD preset logger")
	}

	zapLogger, _ = zap.NewDevelopment()
	zapLogger.Info("Will used DEV preset logger")

	return Logger{zapLogger}
}
