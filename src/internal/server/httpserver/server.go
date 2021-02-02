package httpserver

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Server - ...
type Server struct {
	httpServer *http.Server
}

// Run - ...
func (s *Server) Run(port string, handler http.Handler) {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Println("Server is running on Port: ", port)
}
