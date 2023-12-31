package main

import (
	"gorgom/internal/entity"
	"gorgom/internal/repository"
	"time"
)

func main() {
	db := repository.ConnectDB()

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

	jst, _ := time.LoadLocation("Asia/Tokyo")
	users := []entity.User{
		{
			Email:    "hoge@example.com",
			Name:     "hoge",
			Location: "Tokyo",
			Status:   entity.Free,
		},
		{
			Email:    "fuga@example.com",
			Name:     "fuga",
			Birthday: time.Date(2020, 2, 22, 0, 0, 0, 0, jst),
			Status:   entity.Sick,
		},
	}
	db.Create(&users)

	groups := []entity.Group{{Name: "hogefuga family", Owner: users[0], Members: users}}
	db.Create(&groups)

	roles := []entity.Role{
		{
			Name:    "Mother",
			GroupID: groups[0].ID,
			Users:   []entity.User{users[0]},
		},
		{
			Name:    "Son",
			GroupID: groups[0].ID,
			Users:   []entity.User{users[1]},
		},
	}
	db.Create(&roles)

	for _, user := range users {
		pw := user.Name + user.Name
		user.SetPassword(pw)
		db.Model(&user).Updates(entity.User{Password: user.Password})
	}

	tags := []entity.Tag{
		{Name: "spring"},
		{Name: "summer"},
		{Name: "festival"},
		{Name: "trip"},
	}
	board := entity.Board{Title: "life", OwnerGroup: groups[0], DefinedTags: tags}
	db.Create(&board)

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
		{Title: "food", BoardID: board.ID, Cards: cards[0]},
		{Title: "clothes", BoardID: board.ID, Cards: cards[1]},
		{Title: "house", BoardID: board.ID, Cards: cards[2]},
	}
	db.Create(&boxes)
	db.Model(&board).Updates(entity.Board{Boxes: boxes})
}
