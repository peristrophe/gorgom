package controller

import (
	"gorgom/internal/entity"
	"time"
)

type UserController struct{}

func (uc *UserController) GetUser(r *GetUserRequest) *GetUserResponse {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	user := entity.User{
		Name:      "hoge",
		Status:    entity.Free,
		CreatedAt: now,
		UpdatedAt: now,
	}
	var response GetUserResponse
	response = GetUserResponse(user)
	return &response
}
