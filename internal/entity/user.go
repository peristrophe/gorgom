package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus int

const (
	Free UserStatus = iota
	Busy
	Sick
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string     `json:"name" gorm:"not null"`
	Birthday  *time.Time `json:"birthday"`
	Location  *string    `json:"location"`
	Status    UserStatus `json:"status" gorm:"not null"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt *time.Time `json:"deletedAt"`
}
