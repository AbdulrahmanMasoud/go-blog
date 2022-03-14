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
	//db.Migrator().DropTable(&models.Blog{})
	db.AutoMigrate(&models.User{}, &models.Blog{})

	route := gin.Default()
	route.POST("/register", controllers.Register)
	route.POST("/login", controllers.Login)

	route.Use(middlewares.IsAuth())
	{
		route.GET("/profile", controllers.Profile)

		//Blog Resource
		blogs := route.Group("/blogs")
		{
			blogs.GET("/", controllers.Index)
			blogs.GET("/:id", controllers.Show)
			blogs.POST("/store", controllers.Store)
			blogs.PUT("/:id/update", controllers.Update)
			blogs.DELETE("/:id/delete", controllers.Delete)
		}
	}

	route.Run()

}
