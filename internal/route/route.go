package route

import (
	"gorgom/internal/controller"
	"gorgom/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Route interface{}

type route struct {
	Ctrl controller.Controller
}

func NewRoute(ctrl controller.Controller) *route {
	r := route{Ctrl: ctrl}
	return &r
}

func (r *route) Setup() *gin.Engine {
	gr := gin.Default()

	v1 := gr.Group("/api/v1")
	{
		v1.POST("/signup", r.Ctrl.SignUp())
		v1.POST("/signin", r.Ctrl.SignIn())
	}
	{
		users := v1.Group("/users")
		users.Use(middleware.AuthMiddleware)
		{
			users.GET("/mypage", r.Ctrl.MyPage())
		}

		boards := v1.Group("/boards")
		boards.Use(middleware.AuthMiddleware)
		{
			boards.GET("", r.Ctrl.Boards())
			boards.GET("/:boardID", r.Ctrl.BoardDetail())
		}

		cards := v1.Group("/cards")
		cards.Use(middleware.AuthMiddleware)
		{
			cards.GET("/:cardID", r.Ctrl.CardDetail())
		}
	}
	return gr
}
