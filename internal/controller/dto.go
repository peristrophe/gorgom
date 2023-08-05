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

//type myPageRequest struct {
//	UserID uuid.UUID
//}

type myPageResponse entity.User

type groupDetailRequest struct {
	GroupID uuid.UUID
}

type groupDetailResponse entity.Group

//type groupsRequest struct {}

type groupsResponse []entity.Group

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

func NewBoardDetailRequest(c *gin.Context) *boardDetailRequest {
	boardID := c.Param("boardID")
	boardUUID, err := uuid.Parse(boardID)
	if err != nil {
		return nil
	}
	r := boardDetailRequest{BoardID: boardUUID}
	return &r
}

func NewCardDetailRequest(c *gin.Context) *cardDetailRequest {
	cardID := c.Param("cardID")
	cardUUID, err := uuid.Parse(cardID)
	if err != nil {
		return nil
	}
	r := cardDetailRequest{CardID: cardUUID}
	return &r
}

func NewGroupDetailRequest(c *gin.Context) *groupDetailRequest {
	groupID := c.Param("groupID")
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		return nil
	}
	r := groupDetailRequest{GroupID: groupUUID}
	return &r
}
