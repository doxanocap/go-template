package handler

import (
	"app/internal/manager/interfaces"
	"app/internal/processor/rest/utils"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine       *gin.Engine
	manager      interfaces.IManager
}

func InitHandler(manager interfaces.IManager) *Handler {
	newHandler := &Handler{
		manager: manager,
		Engine: utils.InitEngine(manager.Cfg().App.Environment),
	}

	newHandler.InitRoutes()
	return newHandler
}

func (h *Handler) InitRoutes() {
	h.Engine.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	// here we can add special endpoints
	// based on the Environment
	h.AddRoutesV1()
}
