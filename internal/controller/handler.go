package controller

import (
	"gorgom/internal/entity"
	"gorgom/internal/middleware"
	"gorgom/internal/repository"
	"gorgom/internal/setting"
	"gorgom/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Controller interface {
	SignUp() func(*gin.Context)
	SignIn() func(*gin.Context)
	BoardDetail() func(*gin.Context)
	Boards() func(*gin.Context)
}

type controller struct {
	Repo repository.Repository
}

func NewController(r repository.Repository) *controller {
	ctrl := controller{Repo: r}
	return &ctrl
}

func (ctrl *controller) getAuthorizedUser(c *gin.Context) (*entity.User, error) {
	userID, err := uuid.Parse(c.Param(middleware.CONTEXT_PARAM_KEY_AUTH_USER_ID))
	if err != nil {
		return nil, err
	}

	user, err := ctrl.Repo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
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

		token := util.NewJWT(user.ID.String())
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.SetCookie("token", string(*token), setting.TOKEN_EXPIRE*3600, "/", setting.APP_HOST, false, true)
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
		c.SetCookie("token", string(*token), setting.TOKEN_EXPIRE*3600, "/", setting.APP_HOST, false, true)
		response := signInResponse{UserID: user.ID, Token: token}
		c.IndentedJSON(http.StatusOK, response)
	}
}

func (ctrl *controller) BoardDetail() func(*gin.Context) {
	return func(c *gin.Context) {
		request := NewBoardDetailRequest(c)

		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		board, err := ctrl.Repo.GetBoardByID(request.BoardID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		var groupIDs []uuid.UUID
		for _, group := range user.Groups {
			groupIDs = append(groupIDs, group.ID)
		}
		if !slices.Contains(groupIDs, board.OwnerGroupID) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "board not found."})
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
		}

		var allBoards []entity.Board
		for _, group := range user.Groups {
			boards, err := ctrl.Repo.ListBoardsByGroupID(group.ID)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			allBoards = append(allBoards, boards...)
		}

		response := boardsResponse(allBoards)
		c.IndentedJSON(http.StatusOK, response)
	}
}
