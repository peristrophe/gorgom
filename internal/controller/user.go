package controller

import (
	"gorgom/internal/entity"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

type GetUserRequest struct {
	UserID string
}

func NewGetUserRequest(c *gin.Context) *GetUserRequest {
	uid := c.Param("userID")
	r := GetUserRequest{UserID: uid}
	return &r
}

type GetUserResponse entity.User

func (uc *UserController) GetUser(r *GetUserRequest) *GetUserResponse {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	user := entity.User{
		ID:        r.UserID,
		Name:      "hoge",
		Birthday:  nil,
		Location:  nil,
		Status:    entity.Free,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}
	var response GetUserResponse
	response = GetUserResponse(user)
	return &response
}
