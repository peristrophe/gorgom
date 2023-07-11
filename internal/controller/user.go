package controller

import (
	"gorgom/internal/entity"
	"time"
)

type UserController struct{}

type GetUserRequest struct {
	UserID string
}

func (uc *UserController) GetUser(r *GetUserRequest) *GetUserResponse {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	user := entity.User{
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
