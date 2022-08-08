package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Mycunycu/carsales/gateway/internal/config"
	"github.com/Mycunycu/carsales/gateway/internal/server/httpserver"
	"github.com/Mycunycu/carsales/gateway/pkg/logger"
	"go.uber.org/zap"
)

func Run() error {
	var err error

	err = logger.Init()
	if err != nil {
		return fmt.Errorf("logger.Init %w", err)
	}

	logger := logger.GetLogger()
	defer logger.Sync()
	logger.Info("logger initialized")

	err = config.Init()
	if err != nil {
		return fmt.Errorf("config.Init %w", err)
	}

	cfg := config.GetConfig()
	logger.Info("config initialized")

	server := httpserver.New()

	go func() {
		err := server.Listen(fmt.Sprintf("%s:%s", cfg.HTTPServer.Domain, cfg.HTTPServer.Port))
		if err != nil {
			logger.Error("error while running http server", zap.String("err", err.Error()))
		}
	}()

	logger.Info(fmt.Sprintf("HTTP Server starting on %s:%s", cfg.HTTPServer.Domain, cfg.HTTPServer.Port))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	if err = server.Shutdown(); err != nil {
		logger.Error("failed stopping server", zap.String("err", err.Error()))
	}

	return err
}
