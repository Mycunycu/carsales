package controller

import (
	"carsales/logger"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// SignIn - login
func (auth *AuthController) SignIn(c *fiber.Ctx) error {
	logger := logger.New()
	defer logger.Sync()

	c.SendString("SignIn")
	return nil
}

// SignUp - register
func (auth *AuthController) SignUp(c *fiber.Ctx) error {
	return nil
}
