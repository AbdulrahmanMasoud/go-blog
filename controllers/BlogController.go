package controllers

import (
	"github.com/AbdulrahmanMasoud/blog/database"
	"github.com/AbdulrahmanMasoud/blog/helpers"
	"github.com/AbdulrahmanMasoud/blog/models"
	requests "github.com/AbdulrahmanMasoud/blog/requests/blog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	db := database.Connect()
	var blogs []models.Blog
	db.Where("visible =?", true).Find(&blogs)
	c.JSON(http.StatusOK, gin.H{"data": blogs})
}

// Show blog by id or slug
func Show(c *gin.Context) {
	db := database.Connect()
	var blog models.Blog
	db.Where("id =?", c.Param("id")).Or("slug =?", c.Param("id")).First(&blog)
	if blog.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found this blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": blog})
}

// ShowByUser Show blogs by user
func ShowByUser(c *gin.Context) {
	db := database.Connect()
	var blog []models.Blog
	db.Debug().Where("user_id =?", c.Param("user_id")).Find(&blog)
	//if blog.ID == 0 {
	//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found this blog"})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"data": blog})
}

func Store(c *gin.Context) {
	db := database.Connect()
	var blog requests.CreateBlogRequest
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := helpers.AuthUser(c)
	blog.UserId = user.ID

	created := db.Create(&blog)
	if created.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": created.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": &blog})
}

func Update(c *gin.Context) {
	db := database.Connect()
	var blog requests.UpdateBlogRequest
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := helpers.AuthUser(c)

	updated := db.Where("id =?", c.Param("id")).Where("user_id =?", user.ID).Updates(&blog)

	if updated.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": updated.Error})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "updated done"})
}

func Delete(c *gin.Context) {
	db := database.Connect()
	var blog models.Blog
	user := helpers.AuthUser(c)
	deleted := db.Where("id =?", c.Param("id")).Where("user_id", user.ID).Delete(&blog)
	if deleted.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": deleted.Error})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
