package controllers

import (
	"net/http"
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

	// Parse request body
	var reqBody struct {
		Email    string
		Password string
	}

	c.Bind(&reqBody)

	// Search for user based on email
	var user models.User

	database.DB.Where("email = ?", reqBody.Email).First(&user)

	if user.Id == 0 {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(reqBody.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "incorrect password",
		})
	}

	// JWT

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    user.Id,
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	SecretKey := "0912uiejewfwoefiej"
	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
	}

	// Set cookie content and send it to the client
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

}

func User(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
}

func Logout(c *gin.Context) {
	// JWT

	user, _ := c.Get("user")
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    user.(models.User).Id,
		"ExpiresAt": time.Now().Add(-time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	SecretKey := "0912uiejewfwoefiej"
	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
	}

	// Set cookie content and send it to the client
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, -3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "logged out",
	})

}
