package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser_SetPassword(t *testing.T) {
	salt, _ := uuid.Parse("b665b573-4cd0-4ec8-a52d-dc3ec59df3c0")
	user := User{
		Email:    "hoge@example.com",
		Salt:     salt,
		Name:     "hoge",
		Location: "Tokyo",
	}
	user.SetPassword("hogehoge")
	// depends on env.TOKEN_SECRET_KEY
	assert.Equal(t, "acacb4ae4d947ca15ee0952adaf908012a426e073e413c10285fa412f3135eb3", string(user.Password))
}

func TestUser_Authentication(t *testing.T) {
	salt, _ := uuid.Parse("b665b573-4cd0-4ec8-a52d-dc3ec59df3c0")
	user := User{
		Email:    "hoge@example.com",
		Password: password("acacb4ae4d947ca15ee0952adaf908012a426e073e413c10285fa412f3135eb3"),
		Salt:     salt,
		Name:     "hoge",
		Location: "Tokyo",
	}

	err := user.Authentication("fugafuga")
	assert.NotNil(t, err)

	err = user.Authentication("hogehoge")
	assert.Nil(t, err)
}

func TestUser_ListGroupIDs(t *testing.T) {
	groupID1, _ := uuid.Parse("21f9a09c-b2b5-48cd-b700-80c91b819af9")
	groupID2, _ := uuid.Parse("b69a00d5-54b0-4f41-874d-47e4cbde0256")
	groupID3, _ := uuid.Parse("93cee320-426b-4294-9232-eafddd61ca5b")
	groups := []Group{{ID: groupID1}, {ID: groupID2}, {ID: groupID3}}
	user := User{
		Email: "hoge@example.com",
	}
	assert.Nil(t, user.Groups)
	assert.Empty(t, user.ListGroupIDs())

	user.Groups = make([]Group, 0)
	assert.Empty(t, user.ListGroupIDs())

	user.Groups = groups
	groupIDs := []uuid.UUID{groupID1, groupID2, groupID3}
	assert.Equal(t, groupIDs, user.ListGroupIDs())
}

func TestUser_IsValid(t *testing.T) {
	groupID, _ := uuid.Parse("21f9a09c-b2b5-48cd-b700-80c91b819af9")
	roleID, _ := uuid.Parse("5f411621-ddd3-4568-a9cb-a6d4e54f6ade")
	otherGroupID, _ := uuid.Parse("b69a00d5-54b0-4f41-874d-47e4cbde0256")

	group := Group{ID: groupID}
	role := Role{ID: roleID, GroupID: groupID}
	user := User{
		Email:  "hoge@example.com",
		Groups: []Group{group},
		Roles:  []Role{role},
	}
	assert.True(t, user.IsValid())

	user.Roles[0].GroupID = otherGroupID
	assert.False(t, user.IsValid())
}
