package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IClientService interface {
	NewClient(ctx *gin.Context, conn *websocket.Conn) IClient
}

type IClient interface {
	Reader()
	Writer()
}
