package controller

import (
	"carsales/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct{}

// SignIn - login
func (auth *AuthController) SignIn(ctx *gin.Context) {
	logger := logger.Get()
	defer logger.Sync()

	_, err := ctx.Writer.Write([]byte("SignIn"))
	if err != nil {
		logger.Error("Erorr", zap.NamedError("Err", err))
	}
}

// SignUp - register
func (auth *AuthController) SignUp(ctx *gin.Context) {

}
