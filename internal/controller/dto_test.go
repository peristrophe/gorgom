package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestController_NewBoardDetailRequest(t *testing.T) {
	idStr := "cc6ede1a-c2dc-43e3-a992-ffd8a610be92"
	ctx := gin.Context{}
	ctx.AddParam("boardID", idStr)

	boardID, _ := uuid.Parse(idStr)
	request := NewBoardDetailRequest(&ctx)
	assert.Equal(t, boardID, request.BoardID)
}

func TestController_NewBoardDetailRequest_BadRequest(t *testing.T) {
	ctx := gin.Context{}

	request := NewBoardDetailRequest(&ctx)
	assert.Nil(t, request)
}

func TestController_NewCardDetailRequest(t *testing.T) {
	idStr := "250013d6-6298-4572-932f-ab46dbab0b2c"
	ctx := gin.Context{}
	ctx.AddParam("cardID", idStr)

	cardID, _ := uuid.Parse(idStr)
	request := NewCardDetailRequest(&ctx)
	assert.Equal(t, cardID, request.CardID)
}

func TestController_NewCardDetailRequest_BadRequest(t *testing.T) {
	ctx := gin.Context{}

	request := NewCardDetailRequest(&ctx)
	assert.Nil(t, request)
}

func TestController_NewGroupDetailRequest(t *testing.T) {
	idStr := "40f0e6f9-cc36-49aa-9c73-856c34bcc915"
	ctx := gin.Context{}
	ctx.AddParam("groupID", idStr)

	groupID, _ := uuid.Parse(idStr)
	request := NewGroupDetailRequest(&ctx)
	assert.Equal(t, groupID, request.GroupID)
}

func TestController_NewGroupDetailRequest_BadRequest(t *testing.T) {
	ctx := gin.Context{}

	request := NewGroupDetailRequest(&ctx)
	assert.Nil(t, request)
}
