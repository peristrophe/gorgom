package route

import (
	"gorgom/internal/controller"

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
	}
	{
		boards := v1.Group("/boards")
		{
			//boards.GET("/", ctrl.Boards())
			boards.GET("/:boardID", r.Ctrl.BoardDetail())
		}
	}
	return gr
}
