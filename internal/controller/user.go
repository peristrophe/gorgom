package controller

import (
	"gorgom/internal/setting"
	"gorgom/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *controller) SignUp() func(*gin.Context) {
	return func(c *gin.Context) {
		var request signUpRequest
		if err := c.BindJSON(&request); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := ctrl.Repo.CreateUser(request.Email, request.Password)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		token := util.NewJWT(user.ID.String())
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
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
			return
		}

		user, err := ctrl.Repo.GetUserByEmail(request.Email)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := user.Authentication(request.Password); err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token := util.NewJWT(user.ID.String())
		c.SetCookie("token", string(*token), setting.TOKEN_EXPIRE*3600, "/", setting.APP_HOST, false, true)
		response := signInResponse{UserID: user.ID, Token: token}
		c.IndentedJSON(http.StatusOK, response)
	}
}

func (ctrl *controller) MyPage() func(*gin.Context) {
	return func(c *gin.Context) {
		user, err := ctrl.getAuthorizedUser(c)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if !user.IsValid() {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
			return
		}

		response := myPageResponse(*user)
		c.IndentedJSON(http.StatusOK, response)
	}
}
