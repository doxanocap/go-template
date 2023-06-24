package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine().Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	storage := v1.Group("/storage")
	{
		storage.POST("/:filename", h.ctl.Storage().SaveFile)
		storage.DELETE("/:id")
	}
}
