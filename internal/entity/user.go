package entity

import "time"

type UserStatus int

const (
	Free UserStatus = iota
	Busy
	Sick
)

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Birthday  *time.Time `json:"birthday"`
	Location  *string    `json:"location"`
	Status    UserStatus `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
