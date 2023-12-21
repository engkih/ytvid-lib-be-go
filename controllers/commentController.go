package controllers

import (
	"net/http"

	"example.com/restyt/database"
	"example.com/restyt/models"
	"github.com/gin-gonic/gin"
)

func CommentPost(c *gin.Context) {
	// Catch the req body.
	var reqBody struct {
		VideoId  int
		Username string
		Comment  string
	}

	c.Bind(&reqBody)

	// Put the req body into comment variable.
	comment := models.Comment{VideoId: reqBody.VideoId, Username: reqBody.Username, Comment: reqBody.Comment}
	if comment.VideoId == 0 {
		c.Status(400)
		return
	} else {
		addComment := database.DB.Create(&comment)
		if addComment.Error != nil {
			c.Status(http.StatusBadRequest)
			return
		}
	}

	c.JSON(http.StatusAccepted, gin.H{
		"comment": comment,
	})
}

func CommentIndex(c *gin.Context) {
	var comments []models.Comment
	database.DB.Find(&comments)

	c.JSON(200, gin.H{
		"comments": comments,
	})
}

func CommentUpdate(c *gin.Context) {
	comId := c.Param("comId")

	var reqBody struct {
		VideoId  int
		Username string
		Comment  string
	}
	c.Bind(&reqBody)

	var comment models.Comment
	var comments []models.Comment
	database.DB.First(&comment, comId)

	database.DB.Model(&comment).Updates(models.Comment{VideoId: reqBody.VideoId, Username: reqBody.Username, Comment: reqBody.Comment})

	database.DB.Find(&comments)
	c.JSON(http.StatusAccepted, gin.H{
		"comments": comments,
	})
}

func CommentDelete(c *gin.Context) {
	comId := c.Param("comId")

	var comments []models.Comment

	database.DB.Delete(&models.Comment{}, comId)

	database.DB.Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})

}
