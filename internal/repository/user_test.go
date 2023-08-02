package repository

import (
	"gorgom/internal/entity"
	"gorgom/internal/helper"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	userID, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	salt, _ := uuid.Parse("db00dc13-8698-4a80-9779-76107a82da00")
	rows := sqlmock.NewRows([]string{
		"id",
		"salt",
		"birthday",
		"location",
		"deleted_at",
	}).AddRow(userID, salt, nil, "", nil)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("email","password","name","status","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id","salt","birthday","location","deleted_at"`)).
		WithArgs().
		WillReturnRows(rows)
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "password"=$1,"updated_at"=$2 WHERE "id" = $3`)).
		WithArgs().
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := NewRepository(db)

	user, err := repo.CreateUser("hoge@example.com", "hogehoge")
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

	expect := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Salt:      salt,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	assert.Equal(t, expect, *user)
}

func TestRepository_GetUserByID(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	userID, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	salt, _ := uuid.Parse("5387699f-9cd8-41e0-bb0d-06314803757a")
	rows := sqlmock.NewRows([]string{
		"id",
		"email",
		"password",
		"salt",
		"name",
		"birthday",
		"location",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		userID,
		"hoge@example.com",
		"e109771c585ef37135eacee4835d9229429e0314acb24a9c6b084ab3ec1007b1",
		salt,
		"hoge",
		date,
		"Tokyo",
		entity.Busy,
		date,
		date,
		nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 LIMIT 1`)).
		WithArgs(userID).
		WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "group_users" WHERE "group_users"."user_id" = $1`)).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{}))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "role_users" WHERE "role_users"."user_id" = $1`)).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{}))

	repo := NewRepository(db)

	user, err := repo.GetUserByID(userID)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

	expect := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Salt:      salt,
		Name:      "hoge",
		Birthday:  date,
		Location:  "Tokyo",
		Status:    entity.Busy,
		Groups:    []entity.Group{},
		Roles:     []entity.Role{},
		CreatedAt: date,
		UpdatedAt: date,
	}
	expect.SetPassword("hogehoge")

	assert.Equal(t, expect, *user)
}

func TestRepository_GetUserByEmail(t *testing.T) {
	db, mock, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}
	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	userID, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	salt, _ := uuid.Parse("5387699f-9cd8-41e0-bb0d-06314803757a")
	rows := sqlmock.NewRows([]string{
		"id",
		"email",
		"password",
		"salt",
		"name",
		"birthday",
		"location",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		userID,
		"hoge@example.com",
		"e109771c585ef37135eacee4835d9229429e0314acb24a9c6b084ab3ec1007b1",
		salt,
		"hoge",
		date,
		"Tokyo",
		entity.Busy,
		date,
		date,
		nil,
	)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs("hoge@example.com").
		WillReturnRows(rows)

	repo := NewRepository(db)

	user, err := repo.GetUserByEmail("hoge@example.com")
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}

	expect := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Salt:      salt,
		Name:      "hoge",
		Birthday:  date,
		Location:  "Tokyo",
		Status:    entity.Busy,
		Groups:    nil,
		Roles:     nil,
		CreatedAt: date,
		UpdatedAt: date,
	}
	expect.SetPassword("hogehoge")

	assert.Equal(t, expect, *user)
}
