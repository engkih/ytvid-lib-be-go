package controllers

import (
	"strconv"
	"time"

	"example.com/restyt/database"
	"example.com/restyt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *gin.Context) {

	var reqBody struct {
		Name     string
		Email    string
		Password string
	}

	c.Bind(&reqBody)

	password, _ := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 14)

	user := models.User{Name: reqBody.Name, Email: reqBody.Email, Password: password}

	addUser := database.DB.Create(&user)

	if addUser.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, user)
}

func Login(c *gin.Context) {

	var reqBody struct {
		Name     string
		Email    string
		Password string
	}

	c.Bind(&reqBody)

	var user models.User

	database.DB.Where("email = ?", reqBody.Email).First(&user)

	if user.Id == 0 {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(reqBody.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "incorrect password",
		})
	}

	// JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err = claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.JSON(500, gin.H{
			"message": "could not login",
		})
	}

	//TODO, JWT COOKIE IN GIN-GONIC

	// cookie := map[string]any{
	// 	"Name:":	"jwt",
	// 	"Value":	token,
	// 	"Expires":	time.Now().Add(time.Hour*24),
	// 	"HTTPOnly":	true,
	// }

	// c.Cookie()

}
