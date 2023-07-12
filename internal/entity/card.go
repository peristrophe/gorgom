package entity

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	BoxID       uuid.UUID `json:"boxId"`
	Box         Box       `json:"box"`
	Tags        []Tag     `json:"tags" gorm:"many2many:card_tags"`
	Comments    []Comment `json:"comments" gorm:"foreignkey:CardID"`
	CreatedAt   time.Time `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"default:null"`
}
