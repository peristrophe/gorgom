package handler

import (
	"gorgom/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	request := controller.NewGetUserRequest(c)
	uc := controller.UserController{}
	response := uc.GetUser(request)
	c.JSON(http.StatusOK, response)
}
