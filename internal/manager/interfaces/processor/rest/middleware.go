package rest

import "github.com/gin-gonic/gin"

type IMiddlewareManager interface {
	VerifySession(ctx *gin.Context)
}
