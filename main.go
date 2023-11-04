package main

import (
	//"log"

	"github.com/GingFreecss2/Go-jwt-auth/controllers"
	"github.com/GingFreecss2/Go-jwt-auth/initializers"
	"github.com/GingFreecss2/Go-jwt-auth/middleware"

	//"github.com/GingFreecss2/Go-jwt-auth/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

	// Drop the `users` table.
	//err := initializers.DB.Migrator().DropTable(&models.User{})
	//if err != nil {
	//	log.Printf("Error dropping table: %v", err)
	//}

	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
