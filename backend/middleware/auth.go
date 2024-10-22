package middleware

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Not Authorized"})
		return
	}

	uid, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Not Authorized"})
		return
	}

	c.Set("uid", uid)
	c.Next()
}
