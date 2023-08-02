//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/${GOPACKAGE}_${GOFILE}
package controller

import (
	"gorgom/internal/entity"
	"gorgom/internal/middleware"
	"gorgom/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller interface {
	SignUp() func(*gin.Context)
	SignIn() func(*gin.Context)
	MyPage() func(*gin.Context)

	GroupDetail() func(*gin.Context)
	Groups() func(*gin.Context)

	BoardDetail() func(*gin.Context)
	Boards() func(*gin.Context)

	CardDetail() func(*gin.Context)
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
