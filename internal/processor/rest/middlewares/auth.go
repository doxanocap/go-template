package middlewares

import (
	"app/internal/manager/interfaces"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	manager interfaces.IManager
}

func InitAuthMiddleware(manager interfaces.IManager) *AuthMiddleware {
	return &AuthMiddleware{
		manager: manager,
	}
}

func (am *AuthMiddleware) VerifySession(ctx *gin.Context) {
	//auth := ctx.Request.Header["Authorization"]
	//if len(auth) == 0 {
	//	ctx.JSON(401, models.Error{Status: 401, Message: "User is unauthorized"})
	//	ctx.Abort()
	//	return
	//}
	//
	//auth = strings.Split(auth[0], " ")
	//accessToken := auth[1]
	//user, err := services.ValidateAToken(accessToken)
	//
	//if err.Status != 200 {
	//	ctx.JSON(err.Status, err)
	//	ctx.Abort()
	//	return
	//}
	//
	//ctx.Set("user", user)
	//ctx.Next()
}
