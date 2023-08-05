//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/${GOPACKAGE}_${GOFILE}
package repository

import (
	"gorgom/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(string, string) (*entity.User, error)
	GetUserByID(uuid.UUID) (*entity.User, error)
	GetUserByEmail(string) (*entity.User, error)

	GetBoardByID(uuid.UUID) (*entity.Board, error)
	ListBoardsByGroupID(uuid.UUID) ([]entity.Board, error)

	GetCardByID(uuid.UUID) (*entity.Card, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(defaultDB *gorm.DB) *repository {
	var db *gorm.DB
	if defaultDB == nil {
		//go:nocover
		db = ConnectDB()
	} else {
		db = defaultDB
	}
	repo := repository{DB: db}
	return &repo
}
