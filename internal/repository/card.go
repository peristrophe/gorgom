package repository

import (
	"gorgom/internal/entity"

	"github.com/google/uuid"
)

func (r *repository) GetCardByID(cardID uuid.UUID) (*entity.Card, error) {
	var card entity.Card
	result := r.DB.
		Preload("Tags").
		Preload("Comments").
		Preload("Box.Board").
		Take(&card, cardID)

	if result.Error != nil {
		return nil, result.Error
	}
	return &card, nil
}
