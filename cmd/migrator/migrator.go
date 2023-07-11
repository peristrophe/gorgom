package main

import (
	"gorgom/internal/entity"
	"gorgom/internal/repository"
)

func main() {
	db, err := repository.GetDBConn()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})

	location := "Tokyo"
	user := entity.User{
		Name:     "fuga",
		Location: &location,
		Status:   entity.Free,
	}
	db.Create(&user)
}
