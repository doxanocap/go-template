package handler

import (
	"app/internal/manager/interfaces"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"sync"
)

type Handler struct {
	engine       *gin.Engine
	engineRunner sync.Once
	manager      interfaces.IManager
}

func InitHandler(manager interfaces.IManager) *Handler {
	newHandler := &Handler{
		manager: manager,
	}

	newHandler.InitRoutes()
	return newHandler
}

func (h *Handler) InitRoutes() {
	h.Engine().GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	// here we can add special endpoints
	// based on the Environment
	h.AddRoutesV1()
}

func (h *Handler) Engine() *gin.Engine {
	h.engineRunner.Do(func() {
		h.engine = InitEngine(viper.GetString("ENV_MODE"))
	})
	return h.engine
}

func InitEngine(env string) *gin.Engine {
	setGinMode(env)

	router := gin.New()
	router.Use(gin.Recovery())
	router.RedirectTrailingSlash = true
	corsConfig := cors.Config{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           604800,
	}
	router.Use(cors.New(corsConfig))

	return router
}

func setGinMode(env string) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		return
	}
	gin.SetMode(gin.DebugMode)
}
