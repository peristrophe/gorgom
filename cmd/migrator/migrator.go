package main

import (
	"gorgom/internal/entity"
	"gorgom/internal/repository"
	"time"
)

func main() {
	db, err := repository.GetDBConn()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Role{})
	db.AutoMigrate(&entity.Board{})
	db.AutoMigrate(&entity.Box{})
	db.AutoMigrate(&entity.Card{})
	db.AutoMigrate(&entity.Comment{})
	db.AutoMigrate(&entity.Tag{})
	db.AutoMigrate(&entity.CardTag{})

	roles := []entity.Role{
		{
			Name: "mother",
		},
		{
			Name: "kid",
		},
	}
	jst, err := time.LoadLocation("Asia/Tokyo")
	users := []entity.User{
		{
			Name:     "hoge",
			Location: "Tokyo",
			Status:   entity.Free,
			Role:     roles[0],
		},
		{
			Name:     "fuga",
			Email:    "fuga@example.com",
			Birthday: time.Date(2020, 2, 22, 0, 0, 0, 0, jst),
			Status:   entity.Sick,
			Role:     roles[1],
		},
	}
	db.Create(&users)

	tags := []entity.Tag{
		{Name: "spring"},
		{Name: "summer"},
		{Name: "festival"},
		{Name: "trip"},
	}
	comments := []entity.Comment{
		{Content: "hello, gorgom!", User: users[0]},
		{Content: "good evening", User: users[1]},
		{Content: "my memo", User: users[1]},
	}
	cards := [][]entity.Card{
		{
			{Title: "FOO", Description: "foo", Tags: tags[3:4]},
			{Title: "BAR", Description: "bar"},
			{Title: "BAZ", Description: "baz"},
		},
		{
			{Title: "QUX", Description: "qux", Tags: tags[1:3]},
			{Title: "QUUX", Description: "quux", Comments: comments[:2]},
			{Title: "CORGE", Description: "corge"},
		},
		{
			{Title: "GRAULT", Description: "grault", Comments: comments[2:3], Tags: tags[0:1]},
			{Title: "GARPLY", Description: "garply"},
			{Title: "WALDO", Description: "waldo"},
		},
	}
	boxes := []entity.Box{
		{Title: "food", Cards: cards[0]},
		{Title: "clothes", Cards: cards[1]},
		{Title: "house", Cards: cards[2]},
	}
	board := entity.Board{Title: "life", Boxes: boxes}
	db.Create(&board)
}
