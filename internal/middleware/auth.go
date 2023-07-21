package middleware

import (
	"gorgom/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	token := util.JWT(tokenStr)
	userID := token.WhoAmI()
	if userID == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.AddParam("tokenUserID", userID)
	c.Next()
}
