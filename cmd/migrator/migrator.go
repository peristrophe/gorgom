package main

import (
	"gorgom/internal/entity"
	"gorgom/internal/repository"
	"time"
)

func main() {
	db := repository.GetDBConn()

	db.AutoMigrate(
		&entity.User{},
		&entity.Group{},
		&entity.Role{},
		&entity.Board{},
		&entity.Box{},
		&entity.Card{},
		&entity.Comment{},
		&entity.Tag{},
		&entity.CardTag{},
		&entity.GroupUser{},
	)

	roles := []entity.Role{
		{
			Name: "mother",
		},
		{
			Name: "kid",
		},
	}
	jst, _ := time.LoadLocation("Asia/Tokyo")
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

	groups := []entity.Group{{Owner: users[0], Members: users}}
	db.Create(&groups[0])
	for _, user := range users {
		user.Groups = groups
	}
	db.Update("Groups", &users)

	tags := []entity.Tag{
		{Name: "spring"},
		{Name: "summer"},
		{Name: "festival"},
		{Name: "trip"},
	}
	comments := []entity.Comment{
		{Content: "hello, gorgom!", UserID: users[0].ID},
		{Content: "good evening", UserID: users[1].ID},
		{Content: "my memo", UserID: users[1].ID},
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
	board := entity.Board{Title: "life", OwnerGroup: groups[0], Boxes: boxes}
	db.Create(&board)
}
