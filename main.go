package main

import (
	"github.com/AbdulrahmanMasoud/blog/controllers"
	"github.com/AbdulrahmanMasoud/blog/database"
	"github.com/AbdulrahmanMasoud/blog/middlewares"
	"github.com/AbdulrahmanMasoud/blog/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	conn := database.Connection(db)

	defer conn.Close()

	db.AutoMigrate(&models.User{})

	route := gin.Default()
	route.POST("/register", controllers.Register)
	route.POST("/login", controllers.Login)

	route.Use(middlewares.IsAuth())
	{
		route.GET("/profile", controllers.Profile)
	}

	route.Run()

}
