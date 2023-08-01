//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/${GOPACKAGE}_${GOFILE}
package repository

import (
	"gorgom/internal/entity"

	"github.com/google/uuid"
)

func (r *repository) GetBoardByID(boardID uuid.UUID) (*entity.Board, error) {
	var board entity.Board
	result := r.DB.
		Preload("Boxes.Cards.Comments").
		Preload("Boxes.Cards.Tags").
		Preload("DefinedTags").
		Take(&board, boardID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &board, nil
}

func (r *repository) ListBoardsByGroupID(groupID uuid.UUID) ([]entity.Board, error) {
	var boards []entity.Board
	result := r.DB.Where("owner_group_id = ?", groupID).Find(&boards)
	if result.Error != nil {
		return nil, result.Error
	}
	return boards, nil
}
