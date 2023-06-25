package controllers

import "github.com/gin-gonic/gin"

type IStorageController interface {
	SaveFile(ctx *gin.Context)
}
