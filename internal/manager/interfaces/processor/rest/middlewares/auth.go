package middlewares

import "github.com/gin-gonic/gin"

type IAuthMiddlewares interface {
	VerifySession(ctx *gin.Context)
}
