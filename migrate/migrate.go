package main

import (
	"example.com/restyt/database"
	"example.com/restyt/models"
)

func main() {
	// Connect to datbase.
	database.ConnectDatabase()

	// Create new video Table.
	// database.DB.AutoMigrate(&models.Video{})

	//Create new comment Table.
	database.DB.AutoMigrate(&models.Comment{})
}
