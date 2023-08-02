package main

import (
	"gorgom/internal/controller"
	"gorgom/internal/repository"
	"gorgom/internal/route"
)

func main() {
	repo := repository.NewRepository(nil)
	ctrl := controller.NewController(repo)
	route := route.NewRoute(ctrl)
	r := route.Setup()

	r.Run(":8080")
}
