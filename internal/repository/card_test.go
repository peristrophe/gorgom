package repository

import (
	"fmt"
	"gorgom/internal/entity"
	"gorgom/internal/helper"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetCardByID(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	boxID, _ := uuid.Parse("37cd2a77-1475-4a1d-9a13-b685d4bf6ede")
	cardID, _ := uuid.Parse("250013d6-6298-4572-932f-ab46dbab0b2c")

	cardRows := sqlmock.NewRows([]string{
		"id",
		"title",
		"description",
		"box_id",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		cardID,
		"hoge",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		boxID,
		date,
		date,
		nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cards" WHERE "cards"."id" = $1 LIMIT 1`)).
		WithArgs(cardID).
		WillReturnRows(cardRows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boxes" WHERE "boxes"."id" = $1`)).
		WithArgs(boxID).
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE "comments"."card_id" = $1`)).
		WithArgs(cardID).
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "card_tags" WHERE "card_tags"."card_id" = $1`)).
		WithArgs(cardID).
		WillReturnRows(sqlmock.NewRows([]string{}))
	repo := NewRepository(db)

	card, err := repo.GetCardByID(cardID)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

	expect := entity.Card{
		ID:          cardID,
		Title:       "hoge",
		Description: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		BoxID:       boxID,
		Tags:        []entity.Tag{},
		Comments:    []entity.Comment{},
		CreatedAt:   date,
		UpdatedAt:   date,
	}

	assert.Equal(t, expect, *card)
}

func TestRepository_GetCardByID_SelectError(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	cardID, _ := uuid.Parse("250013d6-6298-4572-932f-ab46dbab0b2c")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cards" WHERE "cards"."id" = $1 LIMIT 1`)).
		WithArgs(cardID).
		WillReturnError(fmt.Errorf("Select failed."))
	repo := NewRepository(db)

	_, err = repo.GetCardByID(cardID)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
