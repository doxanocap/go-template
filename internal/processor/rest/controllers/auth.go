package controllers

import (
	"app/internal/cns/errs"
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
	var body model.SignIn
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := ac.manager.Service().User().Authenticate(ctx, body)
	if err != nil {
		if errs.IsHttpNotFoundError(err) {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, *result)
}

func (ac *AuthController) SignUp(ctx *gin.Context) {
	var body model.SignUp
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := ac.manager.Service().User().Create(ctx, body)
	if err != nil {
		if errs.IsHttpConflictError(err) {
			ctx.Status(http.StatusConflict)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, *result)
}

func (ac *AuthController) Refresh(ctx *gin.Context) {

}

func (ac *AuthController) Logout(ctx *gin.Context) {

}
