package controllers

import (
	"app/internal/manager/interfaces"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	manager interfaces.IManager
}

func InitAuthController(manager interfaces.IManager) *AuthController {
	return &AuthController{
		manager: manager,
	}
}

func (ac *AuthController) SignIn(ctx *gin.Context) {
}

func (ac *AuthController) SignUp(ctx *gin.Context) {
	var body model.SignUp

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

}

func (ac *AuthController) Refresh(ctx *gin.Context) {

}

func (ac *AuthController) Logout(ctx *gin.Context) {

}

func (ac *AuthController) ResetPassword(ctx *gin.Context) {

}

func (ac *AuthController) VerifyEmail(ctx *gin.Context) {

}

func (ac *AuthController) DeleteAccount(ctx *gin.Context) {

}
