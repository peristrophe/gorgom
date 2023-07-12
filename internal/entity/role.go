package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"not null;unique"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt time.Time `json:"deletedAt" gorm:"default:null"`
}
