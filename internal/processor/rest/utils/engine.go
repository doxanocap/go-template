package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
