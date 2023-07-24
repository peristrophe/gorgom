package controller

import (
	"gorgom/internal/entity"
	"gorgom/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type signUpRequest struct {
	Email    string
	Password string
}

type signUpResponse struct {
	UserID uuid.UUID `json:"userID"`
	Token  *util.JWT `json:"token"`
}

type signInRequest struct {
	Email    string
	Password string
}

type signInResponse struct {
	UserID uuid.UUID `json:"userID"`
	Token  *util.JWT `json:"token"`
}

type userProfileRequest struct {
	UserID uuid.UUID
}

type userProfileResponse entity.User

//type boardsRequest struct {
//	GroupID uuid.UUID
//}

type boardsResponse []entity.Board

type boardDetailRequest struct {
	BoardID uuid.UUID
}

type boardDetailResponse entity.Board

type cardDetailRequest struct {
	CardID uuid.UUID
}

type cardDetailResponse entity.Card

func NewUserProfileRequest(c *gin.Context) *userProfileRequest {
	userID := c.Param("userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		panic(err)
	}
	r := userProfileRequest{UserID: userUUID}
	return &r
}

func NewBoardDetailRequest(c *gin.Context) *boardDetailRequest {
	boardID := c.Param("boardID")
	boardUUID, err := uuid.Parse(boardID)
	if err != nil {
		panic(err)
	}
	r := boardDetailRequest{BoardID: boardUUID}
	return &r
}

func NewCardDetailRequest(c *gin.Context) *cardDetailRequest {
	cardID := c.Param("cardID")
	cardUUID, err := uuid.Parse(cardID)
	if err != nil {
		panic(err)
	}
	r := cardDetailRequest{CardID: cardUUID}
	return &r
}
