package controller

import (
	"gorgom/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetUserRequest struct {
	UserID uuid.UUID
}

type GetUserResponse entity.User

type BoardDetailRequest struct {
	BoardID uuid.UUID
}

type BoardDetailResponse entity.Board

func NewGetUserRequest(c *gin.Context) *GetUserRequest {
	userID := c.Param("userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		panic(err)
	}
	r := GetUserRequest{UserID: userUUID}
	return &r
}

func NewBoardDetailRequest(c *gin.Context) *BoardDetailRequest {
	boardID := c.Param("boardID")
	boardUUID, err := uuid.Parse(boardID)
	if err != nil {
		panic(err)
	}
	r := BoardDetailRequest{BoardID: boardUUID}
	return &r
}
