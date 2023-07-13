package entity

import (
	"time"

	"github.com/google/uuid"
)

type CardTag struct {
	CardID    uuid.UUID `json:"cardId" gorm:"primaryKey"`
	TagID     uuid.UUID `json:"tagId" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type GroupUser struct {
	GroupID   uuid.UUID `json:"groupId" gorm:"primaryKey"`
	UserID    uuid.UUID `json:"userId" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}
