package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

func (ctrl *controller) CardDetail() func(*gin.Context) {
	return func(c *gin.Context) {
		request := NewCardDetailRequest(c)

		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		card, err := ctrl.Repo.GetCardByID(request.CardID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		var groupIDs []uuid.UUID
		for _, group := range user.Groups {
			groupIDs = append(groupIDs, group.ID)
		}
		if !slices.Contains(groupIDs, card.Box.Board.OwnerGroupID) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "card not found."})
		}

		response := cardDetailResponse(*card)
		c.IndentedJSON(http.StatusOK, &response)
	}
}
