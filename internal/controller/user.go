package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *controller) MyPage() func(*gin.Context) {
	return func(c *gin.Context) {
		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		response := myPageResponse(*user)
		c.IndentedJSON(http.StatusOK, response)
	}
}
