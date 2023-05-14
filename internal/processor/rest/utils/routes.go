package utils

//package handlers
//
//import (
//	"auth/pkg/configs"
//	"auth/pkg/controllers"
//	"auth/pkg/middlewares"
//	"github.com/gin-contrib/cors"
//	"github.com/gin-gonic/gin"
//)
//
//func SetupRoutes() {
//	port := configs.ENV("PORT")
//	if port == "" {
//		port = "8080"
//	}
//	r := gin.Default()
//	r.Use(cors.New(cors.Config{
//		AllowOrigins:     []string{"http://localhost:3000"},
//		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
//		AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
//		ExposeHeaders:    []string{"Authorization"},
//		AllowCredentials: true,
//	}))
//
//	r.GET("/healthcheck", controllers.Healthcheck)
//
//	auth := r.Group("/auth")
//	{
//		auth.POST("/sign-up", controllers.SignUp)
//		auth.POST("/sign-in", controllers.SignIn)
//		auth.GET("/sign-out", controllers.SignOut)
//		auth.GET("/refresh-token", controllers.RefreshUser)
//
//		user := auth.Group("user")
//		{
//			user.POST("/validate", controllers.AccountInformation)
//
//			user.Use(middlewares.ValidateUserAuth)
//		}
//	}
//
//	r.Run(":" + port)
//}
