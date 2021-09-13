package httpserver

import (
	"carsales/internal/server/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type HttpServer struct {
	*fiber.App
}

func New() *HttpServer {
	app := fiber.New()
	app.Server().MaxConnsPerIP = 1

	routes.Init(app)

	return &HttpServer{app}
}

// Run - ...
func (s *HttpServer) Run(domain string, port string) error {
	return s.Listen(fmt.Sprintf("%s:%s", domain, port))
}

func (s *HttpServer) Stop() error {
	return s.Shutdown()
}
