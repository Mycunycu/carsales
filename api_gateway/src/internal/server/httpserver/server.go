package httpserver

import (
	"carsales/logger"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Run - ...
func Run(port string, handler http.Handler) {
	logger := logger.Get()
	defer logger.Sync()
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logger.Info("Server is running", zap.String("Port", port))
	err := httpServer.ListenAndServe()
	if err != nil {
		logger.Fatal("Serve Error", zap.NamedError("Error", err))
	}
}
