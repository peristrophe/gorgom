package entity

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	return mockDB, mock, err
}

func TestGroup_AfterFind(t *testing.T) {
	db, _, _ := newDBMock()
	groupID, _ := uuid.Parse("4e4d3517-237e-11ee-b7fd-0242ac383802")
	userID0, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	userID1, _ := uuid.Parse("a26427d0-b1d4-489f-8484-19513ac5e732")
	group := Group{
		ID:      groupID,
		Members: []User{{ID: userID0}, {ID: userID1}},
	}
	group.AfterFind(db)

	assert.Equal(t, 2, group.MemberNum)
}
