package main

import (
	"example.com/restyt/controllers"
	"example.com/restyt/database"
	"example.com/restyt/middleware"
	"github.com/gin-gonic/gin"
)

var serverAddrs string = "localhost:8080"

func main() {
	router := gin.Default()
	database.ConnectDatabase()

	router.GET("/api/vidindex", controllers.VideoIndex)
	router.GET("/api/video/:vidId", controllers.VideoShow)
	router.GET("/api/comindex", controllers.CommentIndex)
	router.GET("/api/user", middleware.Authentication, controllers.User)

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
