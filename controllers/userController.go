package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/NerdyNarayan/go-jwt/initializers"
	"github.com/NerdyNarayan/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name" binding:"required"`
	}
	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return

	}
	fmt.Printf("body: %v\n", body)
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{

			"message": "Failed to hash request"})
		return
	}
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not created",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "User created successfully",
	})
	return

}
func SignIn(c *gin.Context) {
	var user models.User
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	initializers.DB.First(&user, "email=?", body.Email)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create token"})
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

}
func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusAccepted, gin.H{"message": user})
}
