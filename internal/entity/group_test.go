package entity

import (
	"gorgom/internal/helper"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGroup_AfterFind(t *testing.T) {
	db, _, err := helper.ConnectMockDB()
	if err != nil {
		panic(err)
	}

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
