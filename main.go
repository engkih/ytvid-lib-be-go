package main

import (
	"time"

	"example.com/restyt/controllers"
	"example.com/restyt/database"
	"example.com/restyt/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var serverAddrs string = "localhost:8080"

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	database.ConnectDatabase()

	router.GET("/api/vidindex", controllers.VideoIndex)
	router.GET("/api/video/:vidId", controllers.VideoShow)
	router.GET("/api/comindex", controllers.CommentIndex)
	router.GET("/api/user", middleware.Authentication, controllers.User)
	router.GET("/api/logout", middleware.Authentication, controllers.Logout)
	router.GET("/api/check", middleware.CookieCheck)

	router.POST("/api/vidpost", controllers.VideoPost)
	router.POST("/api/compost", controllers.CommentPost)
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	router.PUT("/api/video/:vidId", controllers.VideoUpdate)
	router.PUT("/api/comment/:comId", controllers.CommentUpdate)

	router.DELETE("/api/video/:vidId", controllers.VideoDelete)
	router.DELETE("/api/comment/:comId", controllers.CommentDelete)

	router.Run(serverAddrs)
}
