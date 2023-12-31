package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func (ctrl *controller) CardDetail() func(*gin.Context) {
	return func(c *gin.Context) {
		request := NewCardDetailRequest(c)

		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		card, err := ctrl.Repo.GetCardByID(request.CardID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !slices.Contains(user.ListGroupIDs(), card.Box.Board.OwnerGroupID) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Card not found."})
			return
		}

		response := cardDetailResponse(*card)
		c.IndentedJSON(http.StatusOK, &response)
	}
}
