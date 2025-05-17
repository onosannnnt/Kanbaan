package BoardUseCase

import (
	"kanban/entities"
	"kanban/model"
)

type BoardRepository interface {
	Create(board *entities.Board) (*entities.Board, error)
	GetByID(id *string) (*entities.Board, error)
	GetAll() (*[]entities.Board, error)
	Update(board *entities.Board) (*entities.Board, error)
	Delete(id *string) error
	GetByOwnerID(ownerID *string) (*[]entities.Board, error)
	InviteUserToBoard(boardModel *model.InviteUserToBoardInput) (*[]entities.Board, error)
	GetColabBoards(userID *string) (*[]entities.Board, error)
}
