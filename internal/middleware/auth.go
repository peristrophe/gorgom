package middleware

import (
	"gorgom/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const CONTEXT_PARAM_KEY_AUTH_USER_ID = "AuthorizedUserID"

func AuthMiddleware(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	token := util.JWT(tokenStr)
	userID := token.WhoAmI()
	if userID == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.AddParam(CONTEXT_PARAM_KEY_AUTH_USER_ID, userID)
	c.Next()
}
