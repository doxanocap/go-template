package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoutesV1(h *Handler) {
	v1 := h.Engine().Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	storage := v1.Group("/storage")
	{
		//storage.GET("/:id", contro)
		storage.POST("/:key", h.manager.Processor().REST().Controllers().Storage().HandlePicture)
		storage.DELETE("/:id")
	}
}
