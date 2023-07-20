package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"
)

type UserStatus int

const (
	Free UserStatus = iota
	Busy
	Sick
)

type Password string

func (pw Password) Encrypt(salt string) Password {
	pwBytes := []byte(pw)
	saltBytes := []byte(salt)
	cryptedBytes := pbkdf2.Key(pwBytes, saltBytes, 4096, 32, sha256.New)
	crypted := Password(hex.EncodeToString(cryptedBytes))
	return crypted
}

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email     string     `json:"email" gorm:"not null"`
	Password  Password   `json:"-" gorm:"not null"`
	Name      string     `json:"name" gorm:"not null"`
	Birthday  time.Time  `json:"birthday" gorm:"default:null"`
	Location  string     `json:"location" gorm:"default:null"`
	Status    UserStatus `json:"status" gorm:"not null"`
	RoleID    uuid.UUID  `json:"roleId"`
	Role      Role       `json:"role"`
	Groups    []Group    `json:"groups" gorm:"many2many:group_users"`
	CreatedAt time.Time  `json:"createdAt" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"not null;autoUpdateTime"`
	DeletedAt time.Time  `json:"-" gorm:"default:null"`
}
