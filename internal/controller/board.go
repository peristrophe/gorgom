package controller

import (
	"gorgom/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func (ctrl *controller) BoardDetail() func(*gin.Context) {
	return func(c *gin.Context) {
		request := NewBoardDetailRequest(c)

		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		board, err := ctrl.Repo.GetBoardByID(request.BoardID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !slices.Contains(user.ListGroupIDs(), board.OwnerGroupID) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Board not found."})
			return
		}

		response := boardDetailResponse(*board)
		c.IndentedJSON(http.StatusOK, &response)
	}
}

func (ctrl *controller) Boards() func(*gin.Context) {
	return func(c *gin.Context) {
		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var allBoards []entity.Board
		for _, group := range user.Groups {
			boards, err := ctrl.Repo.ListBoardsByGroupID(group.ID)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			allBoards = append(allBoards, boards...)
		}

		response := boardsResponse(allBoards)
		c.IndentedJSON(http.StatusOK, response)
	}
}
