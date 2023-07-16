//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/${GOPACKAGE}_${GOFILE}
package repository

import (
	"gorgom/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	BoardByID(uuid.UUID) *entity.Board
	BoardsByGroupID(uuid.UUID) []*entity.Board
}

type repository struct {
	DB *gorm.DB
}

func NewRepository() *repository {
	db := GetDBConn()
	repo := repository{DB: db}
	return &repo
}

func (r *repository) BoardByID(boardId uuid.UUID) *entity.Board {
	var board entity.Board
	r.DB.
		Preload("Boxes.Cards.Comments").
		Preload("Boxes.Cards.Tags").
		Take(&board, boardId)

	return &board
}

func (r *repository) BoardsByGroupID(groupId uuid.UUID) []*entity.Board {
	return nil
}
