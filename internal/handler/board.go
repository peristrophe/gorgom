package handler

import (
	"gorgom/internal/controller"
	"gorgom/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BoardDetail(c *gin.Context) {
	request := controller.NewBoardDetailRequest(c)
	repo := repository.NewBoardRepository()
	board := repo.BoardByID(request.BoardID)
	c.IndentedJSON(http.StatusOK, board)
}
