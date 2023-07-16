package main

import (
	"gorgom/internal/controller"
	"gorgom/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := repository.NewRepository()
	ctrl := controller.NewController(repo)

	v1 := r.Group("/v1")
	{
		boards := v1.Group("/boards")
		{
			//boards.GET("/", ctrl.Boards())
			boards.GET("/:boardID", ctrl.BoardDetail())
		}
	}

	r.Run(":8080")
}
