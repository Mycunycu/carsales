package httpserver

import (
	"carsales/internal/server/routes"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	*fiber.App
}

func New() *HTTPServer {
	app := fiber.New()
	app.Server().MaxConnsPerIP = 1

	routes.Init(app)

	return &HTTPServer{app}
}
