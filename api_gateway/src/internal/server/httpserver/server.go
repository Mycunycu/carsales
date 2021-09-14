package httpserver

import (
	"carsales/internal/server/routes"

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
