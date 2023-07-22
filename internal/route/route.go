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
		boards := v1.Group("/boards")
		boards.Use(middleware.AuthMiddleware)
		{
			boards.GET("/", r.Ctrl.Boards())
			boards.GET("/:boardID", r.Ctrl.BoardDetail())
		}
	}
	return gr
}
