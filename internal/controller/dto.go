package controller

import (
	"gorgom/internal/entity"

	"github.com/gin-gonic/gin"
)

func NewGetUserRequest(c *gin.Context) *GetUserRequest {
	uid := c.Param("userID")
	r := GetUserRequest{UserID: uid}
	return &r
}

type GetUserResponse entity.User
