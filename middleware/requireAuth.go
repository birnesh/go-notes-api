package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/birnesh/go-notes-api/initializers"
	"github.com/birnesh/go-notes-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// get the JWT token
	var tokenVariable = "Authorization"
	tokenString := c.GetHeader(tokenVariable)
	if len(tokenString) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	tokenString = tokenString[len("Bearer "):]
	// decode/validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check the token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		c.Next()
	} else {
		fmt.Println(err)
	}

}
