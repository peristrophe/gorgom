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

func (ctrl *controller) SignUp() func(*gin.Context) {
	return func(c *gin.Context) {
		var request signUpRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		if err := ctrl.Repo.CreateUser(request.Email, request.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}

func (ctrl *controller) SignIn() func(*gin.Context) {
	return func(c *gin.Context) {
		var request signInRequest
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
		token, err := ctrl.Repo.Authentication(request.Email, request.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
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
