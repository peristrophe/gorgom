package repository

import (
	"gorgom/internal/entity"

	"github.com/google/uuid"
)

type BoardRepository interface {
	BoardByID(uuid.UUID) *entity.Board
	BoardsByGroupID(uuid.UUID) []*entity.Board
}

type boardRepository struct{}

func NewBoardRepository() BoardRepository {
	br := boardRepository{}
	return &br
}

func (br *boardRepository) BoardByID(boardId uuid.UUID) *entity.Board {
	db, err := GetDBConn()
	if err != nil {
		return nil
	}

	var board entity.Board
	db.Preload("Boxes.Cards.Comments").Preload("Boxes.Cards.Tags").Preload("OwnerGroup.Owner.Role").Preload("OwnerGroup.Members.Role").Take(&board, boardId)
	return &board
}

func (be *boardRepository) BoardsByGroupID(groupId uuid.UUID) []*entity.Board {
	return nil
}
