package main

import (
	"fmt"
	"github.com/NerdyNarayan/go-jwt/controllers"
	"github.com/NerdyNarayan/go-jwt/initializers"
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
	r.Run(":8001")
}
