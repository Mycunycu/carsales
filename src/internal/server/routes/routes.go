package routes

import (
	"carsales/internal/controller"
	"carsales/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Get() *gin.Engine {
	r := gin.Default()

	corsMiddleware := new(middleware.Cors).Init()
	r.Use(corsMiddleware)

	authController := new(controller.AuthController)
	auth := r.Group("api/auth")
	{
		auth.POST("/signUp", authController.SignUp)
		auth.POST("/signIn", authController.SignIn)
	}

	return r
}
