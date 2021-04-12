package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthController struct{}

// SignIn - login
func (auth *AuthController) SignIn(ctx *gin.Context) {

	_, err := ctx.Writer.Write([]byte("SignIn"))
	if err != nil {
		logrus.Println(err)
	}
}

// SignUp - register
func (auth *AuthController) SignUp(ctx *gin.Context) {

}
