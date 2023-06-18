package controllers

import "github.com/gin-gonic/gin"

type IStorageManager interface {
	HandlePicture(ctx *gin.Context)
}
