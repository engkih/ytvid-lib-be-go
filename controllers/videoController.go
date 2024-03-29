package controllers

import (
	"net/http"

	"example.com/restyt/database"
	"example.com/restyt/models"
	"github.com/gin-gonic/gin"
)

func VideoPost(c *gin.Context) {

	// Pull req body and put it into "video" variable.
	var reqBody struct {
		UserId      int
		VideoUrl    string
		Title       string
		Description string
	}

	c.Bind(&reqBody)

	video := models.Video{UserId: reqBody.UserId, VideoUrl: reqBody.VideoUrl, Title: reqBody.Title, Description: reqBody.Description}

	// Add video to the database.
	addVideo := database.DB.Create(&video)

	if addVideo.Error != nil {
		c.Status(400)
		return
	}

	// Send http respond.
	c.JSON(201, gin.H{
		"video": video,
	})
}

func VideoIndex(c *gin.Context) {
	user, _ := c.Get("user")

	userId := user.(models.User).Id

	var videos []models.Video
	// database.DB.Find(&videos)
	// database.DB.Where("UserId <> ?", ids).Find(&videos)
	database.DB.Where("user_id = ?", userId).Find(&videos)
	// fmt.Println(userId)

	c.JSON(200, gin.H{
		"User":   user,
		"videos": videos,
	})
}

func VideoShow(c *gin.Context) {

	vidId := c.Param("vidId")
	user, _ := c.Get("user")

	var video models.Video

	database.DB.First(&video, vidId)

	c.JSON(http.StatusAccepted, gin.H{
		"user":  user,
		"video": video,
		"test":  "abcdefg",
	})
}

func VideoUpdate(c *gin.Context) {
	// Catch req data.
	vidId := c.Param("vidId")

	var reqBody struct {
		VideoUrl    string
		Title       string
		Description string
	}
	c.Bind(&reqBody)

	// Search for the data that want to be edited.
	var video models.Video
	database.DB.First(&video, vidId)

	// Edit the data that has been searched.
	database.DB.Model(&video).Updates(models.Video{
		VideoUrl: reqBody.VideoUrl, Title: reqBody.Title, Description: reqBody.Description,
	})

	// Respond succes.
	c.JSON(202, gin.H{
		"video": video,
	})
}

func VideoDelete(c *gin.Context) {
	vidId := c.Param("vidId")

	database.DB.Delete(&models.Video{}, vidId)

	c.Status(200)
}
