package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CardTag struct {
	CardID    uuid.UUID `json:"cardId" gorm:"primaryKey"`
	TagID     uuid.UUID `json:"tagId" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

func (ct *CardTag) BeforeCreate(db *gorm.DB) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	ct.CreatedAt = time.Now().In(jst)
	return nil
}
