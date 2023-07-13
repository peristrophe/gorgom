package controller

import "gorgom/internal/repository"

type BoardController struct{}

func (bc *BoardController) GetBoardDetail(r *BoardDetailRequest) *BoardDetailResponse {
	// TODO permission control
	repo := repository.NewBoardRepository()
	board := repo.BoardByID(r.BoardID)
	response := BoardDetailResponse(*board)
	return &response
}
