package helpers

import (
	"github.com/AbdulrahmanMasoud/blog/database"
	"github.com/AbdulrahmanMasoud/blog/models"
	"github.com/AbdulrahmanMasoud/blog/token"
	"github.com/gin-gonic/gin"
)

// AuthUser To return authentication user
func AuthUser(c *gin.Context) models.User {
	db := database.Connect()
	var user models.User
	userId, _ := token.ExtractTokenID(c)
	db.Find(&user, userId)
	return user
}
