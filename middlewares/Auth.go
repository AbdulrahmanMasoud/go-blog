package middlewares

import (
	"github.com/AbdulrahmanMasoud/blog/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
