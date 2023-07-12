package entity

import (
	"time"

	"github.com/google/uuid"
)

type Box struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title     string    `json:"title" gorm:"not null"`
	BoardID   uuid.UUID `json:"boardId"`
	Board     Board     `json:"board"`
	Cards     []Card    `json:"cards" gorm:"foreignkey:BoxID"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt time.Time `json:"deletedAt" gorm:"default:null"`
}
