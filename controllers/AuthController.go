package controllers

import (
	"github.com/AbdulrahmanMasoud/blog/database"
	"github.com/AbdulrahmanMasoud/blog/models"
	"github.com/AbdulrahmanMasoud/blog/requests"
	"github.com/AbdulrahmanMasoud/blog/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	db := database.Connect()
	var user requests.RegisterUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
}

func Login(c *gin.Context) {

	db := database.Connect()

	var user requests.LoginUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//To ge binded password
	password := user.Password
	db.Where("email = ?", user.Email).First(&user)
	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email Or Password is wrong"})
		return
	}

	err := user.ComparePassword(password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email Or Password is wrong"})
		return
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"token":   token,
		"message": "You are loge in",
	})

}

func Profile(c *gin.Context) {
	db := database.Connect()
	var user models.User
	userId, _ := token.ExtractTokenID(c)

	db.Find(&user, userId)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
