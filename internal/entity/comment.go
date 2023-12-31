package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uuid.UUID `json:"userId"`
	User      User      `json:"-"`
	CardID    uuid.UUID `json:"cardId"`
	Card      Card      `json:"-"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt time.Time `json:"-" gorm:"default:null"`
}
