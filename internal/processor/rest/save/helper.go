package save

//package handlers
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"github.com/sirupsen/logrus"
//	"net/http"
//	"webchat/pkg/repository/models"
//)
//
//var ConnUpgrade = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//	CheckOrigin:     func(r *http.Request) bool { return true },
//}
//
//func (handler *Handler) newErrorResponse(ctx *gin.Context, err models.ErrorResponse) {
//	logrus.Error(err.Message)
//	ctx.AbortWithStatusJSON(err.Status, gin.H{"status": err.Status, "message": err.Message})
//}
//
//func (handler *Handler) healthcheck(ctx *gin.Context) {
//	ctx.JSON(200, "HANDLER service is alive")
//}
