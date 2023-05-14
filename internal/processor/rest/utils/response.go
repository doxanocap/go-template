package utils

//package handlers
//
//import (
//	"api/pkg/repository"
//
//	"github.com/gin-gonic/gin"
//	"github.com/sirupsen/logrus"
//)
//
//func (router *Router) newErrorResponse(ctx *gin.Context, err repository.ErrorResponse) {
//	logrus.Error(err.Message)
//	ctx.AbortWithStatusJSON(err.Status, gin.H{"status": err.Status, "message": err.Message})
//}
//
//func (router *Router) healthcheck(ctx *gin.Context) {
//	ctx.JSON(200, gin.H{"message": "router service is alive"})
//}
