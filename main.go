package main

import (
	"github.com/NerdyNarayan/go-jwt/controllers"
	"github.com/NerdyNarayan/go-jwt/initializers"
	"github.com/NerdyNarayan/go-jwt/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.GET("/validate", middlewares.AuthMiddlerware, controllers.Validate)
	r.Run(":8001")
}
