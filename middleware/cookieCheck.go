package middleware

import (
	"fmt"
	"net/http"
	"time"

	"example.com/restyt/database"
	"example.com/restyt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// const SecretKey = "0912uiejewfwoefiej"

func CookieCheck(c *gin.Context) {

	// Get cookie from client.
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}
	// Parse JWT.
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok {

		//Check expiration date.
		if float64(time.Now().Unix()) > claims["ExpiresAt"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
		}

		// Find user based on Id.
		var user models.User
		database.DB.First(&user, claims["Issuer"])

		if user.Id == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
		}

		c.JSON(http.StatusAccepted, gin.H{
			"UserDat": user,
		})

		// fmt.Println(claims["Issuer"], claims["ExpiresAt"])
	} else {
		fmt.Println(err)
	}

}
