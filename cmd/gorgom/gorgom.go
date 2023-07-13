package main

import (
	"gorgom/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/:userID", handler.GetUser)

	r.GET("/board/:boardID", handler.BoardDetail)

	r.Run(":8080")
}
