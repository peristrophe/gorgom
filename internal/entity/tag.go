package entity

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"content" gorm:"not null"`
	BoardID   uuid.UUID `json:"boardId"`
	Cards     []Card    `json:"-" gorm:"many2many:card_tags"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt time.Time `json:"-" gorm:"default:null"`
}
