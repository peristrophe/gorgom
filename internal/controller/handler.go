package controller

import (
	"gorgom/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	BoardDetail() func(*gin.Context)
}

type controller struct {
	Repo repository.Repository
}

func NewController(r repository.Repository) *controller {
	ctrl := controller{Repo: r}
	return &ctrl
}

func (ctrl *controller) BoardDetail() func(*gin.Context) {
	return func(c *gin.Context) {
		// TODO permission control
		request := NewBoardDetailRequest(c)
		board := ctrl.Repo.BoardByID(request.BoardID)
		response := boardDetailResponse(*board)
		c.IndentedJSON(http.StatusOK, &response)
	}
}
