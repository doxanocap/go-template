package handler

import (
	"app/internal/manager/interfaces"
	"app/internal/model"
	"app/internal/processor/rest/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	engine *gin.Engine

	log     *zap.Logger
	cfg     *model.Config
	manager interfaces.IManager
}

func InitHandler(manager interfaces.IManager, log *zap.Logger) *Handler {
	newHandler := &Handler{
		log:     log,
		manager: manager,
		cfg:     manager.Cfg(),
	}

	newHandler.engine = utils.InitEngine(newHandler.cfg.Environment)
	newHandler.InitRoutes()
	return newHandler
}

func (h *Handler) Engine() *gin.Engine {
	return h.engine
}

func (h *Handler) InitRoutes() {
	h.Engine().GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	// here we can add special endpoints
	// based on the Environment
	h.AddRoutesV1()
}
