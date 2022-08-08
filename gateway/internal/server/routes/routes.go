package routes

import (
	"github.com/Mycunycu/carsales/gateway/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("version", "v1")

		return c.Next()
	})

	authController := controller.NewAuthController()
	v1.Post("/signIn", authController.SignIn)
	v1.Post("/signUp", authController.SignUp)
}
