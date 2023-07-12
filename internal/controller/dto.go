package controller

import (
	"gorgom/internal/entity"

	"github.com/gin-gonic/gin"
)

type GetUserRequest struct {
	UserID string
}

type GetUserResponse entity.User

func NewGetUserRequest(c *gin.Context) *GetUserRequest {
	uid := c.Param("userID")
	r := GetUserRequest{UserID: uid}
	return &r
}
