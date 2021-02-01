package routes

import (
	"carsales/internal/controller"
	"carsales/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// Routes ...
type Routes struct{}

// Init ...
func (r *Routes) Init() *gin.Engine {
	router := gin.Default()

	corsMiddleware := new(middleware.Cors).Init()
	router.Use(corsMiddleware)

	authController := new(controller.AuthController)
	auth := router.Group("/auth")
	{
		auth.POST("/signUp", authController.SignUp)
		auth.POST("/signIn", authController.SignIn)
	}

	return router
}
