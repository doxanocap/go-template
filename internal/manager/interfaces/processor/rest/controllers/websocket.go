package controllers

import "github.com/gin-gonic/gin"

type IWSController interface {
	Pool(ctx *gin.Context)
}
