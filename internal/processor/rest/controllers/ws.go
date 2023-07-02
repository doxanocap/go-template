package controllers

import (
	"app/internal/manager/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebsocketController struct {
	manager interfaces.IManager
}

func InitWebsocketController(manager interfaces.IManager) *WebsocketController {
	return &WebsocketController{
		manager: manager,
	}
}

var connectionUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (ws *WebsocketController) Pool(ctx *gin.Context) {
	header := http.Header{}
	conn, err := connectionUpgrade.Upgrade(ctx.Writer, ctx.Request, header)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ws.manager.Processor().WS().Pool()
	client := ws.manager.Processor().WS().Client().NewClient(ctx, conn)

	go client.Reader()
	go client.Writer()
}
