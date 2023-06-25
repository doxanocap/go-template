package controllers

import "github.com/gin-gonic/gin"

type IAuthController interface {
	SignIn(ctx *gin.Context)
	SignUp(ctx *gin.Context)
	Refresh(ctx *gin.Context)
	Logout(ctx *gin.Context)
}
