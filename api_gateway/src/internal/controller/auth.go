package controller

import (
	"carsales/logger"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct{}

// SignIn - login
func (auth *AuthController) SignIn(ctx *fiber.Ctx) error {
	logger := logger.Get()
	defer logger.Sync()

	ctx.SendString("SignIn")
	return nil
}

// SignUp - register
func (auth *AuthController) SignUp(ctx *fiber.Ctx) error {
	return nil
}
