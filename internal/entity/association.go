package entity

import (
	"github.com/google/uuid"
)

type CardTag struct {
	CardID uuid.UUID `json:"cardId" gorm:"primaryKey"`
	TagID  uuid.UUID `json:"tagId" gorm:"primaryKey"`
}

type GroupUser struct {
	GroupID uuid.UUID `json:"groupId" gorm:"primaryKey"`
	UserID  uuid.UUID `json:"userId" gorm:"primaryKey"`
}

type RoleUser struct {
	RoleID uuid.UUID `json:"roleId" gorm:"primaryKey"`
	UserID uuid.UUID `json:"userId" gorm:"primaryKey"`
}
