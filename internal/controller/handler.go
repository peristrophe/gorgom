package controller

import (
	"fmt"
	"gorgom/internal/repository"
	"gorgom/internal/setting"
	"gorgom/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	SignUp() func(*gin.Context)
	SignIn() func(*gin.Context)
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
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			panic(err)
		}
		fmt.Printf("%v", request)
		user, err := ctrl.Repo.CreateUser(request.Email, request.Password)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}

		token, err := util.GenerateToken(user.ID.String())
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.SetCookie("token", token, setting.TOKEN_EXPIRE*3600, "/", "localhost", false, true)
		response := signUpResponse{UserID: user.ID, Token: token}
		c.IndentedJSON(http.StatusOK, response)
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
