package database

import (
	"log"

	"example.com/restyt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("engkih:123321@tcp(localhost:3306)/ytlibdb_go?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	log.Println("Database connected.")
	DB = database

	database.AutoMigrate(&models.Video{})
	database.AutoMigrate(&models.User{})
}
