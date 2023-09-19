package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var connectionUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *Handler) Pool(ctx *gin.Context) {
	header := http.Header{}
	conn, err := connectionUpgrade.Upgrade(ctx.Writer, ctx.Request, header)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.manager.Processor().WS().Pool()
	client := h.manager.Processor().WS().Client().NewClient(ctx, conn)

	go client.Reader()
	go client.Writer()
}
