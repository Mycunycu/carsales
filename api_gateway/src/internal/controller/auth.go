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

	logger.Info("Hey")
	return nil
}

// SignUp - register
func (auth *AuthController) SignUp(ctx *fiber.Ctx) error {
	return nil
}
