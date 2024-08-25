package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/NerdyNarayan/go-jwt/initializers"
	"github.com/NerdyNarayan/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddlerware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var user models.User
		initializers.DB.First(&user, "id=?", claims["sub"])
		if user.ID == 0 {

			c.AbortWithStatus(http.StatusUnauthorized)
			return

		}
		c.Set("user", user)
	} else {

		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
