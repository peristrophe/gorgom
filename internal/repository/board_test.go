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

func TestRepository_ListBoardsByGroupID(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	boardID, _ := uuid.Parse("cc6ede1a-c2dc-43e3-a992-ffd8a610be92")
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")

	rows := sqlmock.NewRows([]string{
		"id",
		"title",
		"owner_group_id",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		boardID,
		"hoge",
		groupID,
		date,
		date,
		nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE owner_group_id = $1`)).
		WithArgs(groupID).
		WillReturnRows(rows)
	repo := NewRepository(db)

	boards, err := repo.ListBoardsByGroupID(groupID)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

	expect := []entity.Board{
		{
			ID:           boardID,
			Title:        "hoge",
			OwnerGroupID: groupID,
			CreatedAt:    date,
			UpdatedAt:    date,
		},
	}

	assert.Equal(t, expect, boards)
}

func TestRepository_ListBoardsByGroupID_SelectError(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE owner_group_id = $1`)).
		WithArgs(groupID).
		WillReturnError(fmt.Errorf("Select failed."))
	repo := NewRepository(db)

	_, err = repo.ListBoardsByGroupID(groupID)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestRepository_GetBoardByID(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	boardID, _ := uuid.Parse("cc6ede1a-c2dc-43e3-a992-ffd8a610be92")
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")

	boardRows := sqlmock.NewRows([]string{
		"id",
		"title",
		"owner_group_id",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		boardID,
		"hoge",
		groupID,
		date,
		date,
		nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE "boards"."id" = $1 LIMIT 1`)).
		WithArgs(boardID).
		WillReturnRows(boardRows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boxes" WHERE "boxes"."board_id" = $1`)).
		WithArgs(boardID).
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tags" WHERE "tags"."board_id" = $1`)).
		WithArgs(boardID).
		WillReturnRows(sqlmock.NewRows([]string{}))
	repo := NewRepository(db)

	board, err := repo.GetBoardByID(boardID)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

	expect := entity.Board{
		ID:           boardID,
		Title:        "hoge",
		OwnerGroupID: groupID,
		Boxes:        []entity.Box{},
		DefinedTags:  []entity.Tag{},
		CreatedAt:    date,
		UpdatedAt:    date,
	}

	assert.Equal(t, expect, *board)
}

func TestRepository_GetBoardByID_SelectError(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	boardID, _ := uuid.Parse("cc6ede1a-c2dc-43e3-a992-ffd8a610be92")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "boards" WHERE "boards"."id" = $1 LIMIT 1`)).
		WithArgs(boardID).
		WillReturnError(fmt.Errorf("Select failed."))
	repo := NewRepository(db)

	_, err = repo.GetBoardByID(boardID)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
