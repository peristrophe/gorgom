package entity

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title        string    `json:"title" gorm:"not null"`
	OwnerGroupID uuid.UUID `json:"ownerGroupId"`
	OwnerGroup   Group     `json:"-"`
	Boxes        []Box     `json:"boxes" gorm:"foreignkey:BoardID"`
	DefinedTags  []Tag     `json:"definedTags" gorm:"foreignkey:BoardID"`
	CreatedAt    time.Time `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt    time.Time `json:"-" gorm:"default:null"`
}
