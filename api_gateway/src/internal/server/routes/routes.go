package routes

import (
	"carsales/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1", func(ctx *fiber.Ctx) error {
		ctx.Set("version", "v1")
		return ctx.Next()
	})

	authController := controller.NewAuthController()
	v1.Post("/signIn", authController.SignIn)
	v1.Post("/signUp", authController.SignUp)
}
