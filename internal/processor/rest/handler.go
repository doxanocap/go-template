package rest

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Engine *gin.Engine
}

func InitHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitEngine(env string) {
	setGinMode(env)

	router := gin.New()
	router.Use(gin.Recovery())
	// todo: cors from app config
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

	h.Engine = router
}

func (h *Handler) AddRoutes() {
	v1 := h.Engine.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

}

func setGinMode(env string) {
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
		return
	}
	gin.SetMode(gin.DebugMode)
}
