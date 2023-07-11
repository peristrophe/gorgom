package main

import (
	"gorgom/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/:userID", handler.GetUser)

	r.Run(":8080")
}
