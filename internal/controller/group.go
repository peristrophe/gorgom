package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *controller) GroupDetail() func(*gin.Context) {
	return func(c *gin.Context) {
		request := NewGroupDetailRequest(c)

		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for _, group := range user.Groups {
			if group.ID == request.GroupID {
				response := groupDetailResponse(group)
				c.IndentedJSON(http.StatusOK, response)
				return
			}
		}

		msg := fmt.Sprintf("Not found group: %s", request.GroupID.String())
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": msg})
	}
}

func (ctrl *controller) Groups() func(*gin.Context) {
	return func(c *gin.Context) {
		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		size := len(user.Groups)
		for i := 0; i < size; i++ {
			user.Groups[i].Members = nil
		}

		response := groupsResponse(user.Groups)
		c.IndentedJSON(http.StatusOK, response)
	}
}
