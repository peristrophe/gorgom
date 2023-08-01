package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string    `json:"name"`
	OwnerID      uuid.UUID `json:"ownerId"`
	Owner        User      `json:"-"`
	MemberNum    int       `json:"memberNum" gorm:"-"`
	Members      []User    `json:"members" gorm:"many2many:group_users"`
	DefinedRoles []Role    `json:"definedRoles" gorm:"foreignkey:GroupID"`
	CreatedAt    time.Time `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt    time.Time `json:"-" gorm:"default:null"`
}

func (g *Group) AfterFind(tx *gorm.DB) error {
	if g.Members != nil {
		(*g).MemberNum = len((*g).Members)
	}
	return nil
}
