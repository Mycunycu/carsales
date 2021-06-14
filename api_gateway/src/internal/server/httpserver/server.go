package httpserver

import (
	"carsales/internal/server/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Run - ...
func Run(domain string, port string) error {
	app := fiber.New()
	app.Server().MaxConnsPerIP = 1

	routes.Init(app)

	if err := app.Listen(fmt.Sprintf("%s:%s", domain, port)); err != nil {
		return err
	}

	return nil
}
