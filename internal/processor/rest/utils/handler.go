package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Engine *gin.Engine
}

func InitHandler() *Handler {
	return &Handler{
		Engine: InitEngine("env"),
	}
}

func (h *Handler) AddRoutesV1() {
	v1 := h.Engine.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
}
