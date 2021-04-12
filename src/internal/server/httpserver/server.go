package httpserver

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Run - ...
func Run(port string, handler http.Handler) {
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logrus.Println("Server is running on Port: ", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}
