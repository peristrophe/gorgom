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

		user, err := ctrl.Repo.CreateUser(request.Email, request.Password)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			panic(err)
		}

		//token, err := util.GenerateToken(user.ID.String())
		token := util.NewJWT(user.ID.String())
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.SetCookie("token", string(*token), setting.TOKEN_EXPIRE*3600, "/", "localhost", false, true)
		response := signUpResponse{UserID: user.ID, Token: token}
		c.IndentedJSON(http.StatusOK, response)
	}
}

func (ctrl *controller) SignIn() func(*gin.Context) {
	return func(c *gin.Context) {
		var request signInRequest
		if err := c.BindJSON(&request); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user, err := ctrl.Repo.GetUserByEmail(request.Email)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if err := user.Authentication(request.Password); err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		token := util.NewJWT(user.ID.String())
		c.SetCookie("token", string(*token), setting.TOKEN_EXPIRE*3600, "/", "localhost", false, true)
		response := signInResponse{UserID: user.ID, Token: token}
		c.JSON(http.StatusOK, response)
	}
}

func (ctrl *controller) BoardDetail() func(*gin.Context) {
	return func(c *gin.Context) {
		// TODO permission control
		var uid string
		uid = c.Param("tokenUserID")
		fmt.Println(uid)

		request := NewBoardDetailRequest(c)
		board := ctrl.Repo.BoardByID(request.BoardID)
		response := boardDetailResponse(*board)
		c.IndentedJSON(http.StatusOK, &response)
	}
}
