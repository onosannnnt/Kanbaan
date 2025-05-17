package BoardUseCase

import (
	"errors"
	"kanban/entities"
	"kanban/model"
)

type BoardUseCase interface {
	Create(board *entities.Board) (*entities.Board, error)
	GetByID(id *string) (*entities.Board, error)
	GetAll() (*[]entities.Board, error)
	Update(board *entities.Board) (*entities.Board, error)
	Delete(id *string) error
	InviteUserToBoard(boardModel *model.InviteUserToBoardInput) (*[]entities.Board, error)
}

type BoardService struct {
	BoardRepository BoardRepository
}

func NewBoardUseCase(boardRepository BoardRepository) BoardUseCase {
	return &BoardService{
		BoardRepository: boardRepository,
	}
}

func (b *BoardService) Create(board *entities.Board) (*entities.Board, error) {
	createdBoard, err := b.BoardRepository.Create(board)
	if err != nil {
		return nil, err
	}
	return createdBoard, nil
}

func (b *BoardService) GetByID(id *string) (*entities.Board, error) {
	board, err := b.BoardRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (b *BoardService) GetAll() (*[]entities.Board, error) {
	boards, err := b.BoardRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return boards, nil
}

func (b *BoardService) Update(board *entities.Board) (*entities.Board, error) {
	boardID := board.ID.String()
	selectedBoard, err := b.BoardRepository.GetByID(&boardID)
	if err != nil {
		return nil, err
	}
	if selectedBoard == nil {
		return nil, err
	}
	if selectedBoard.OwnerID != board.OwnerID {
		return nil, err
	}
	if board.Name != "" {
		selectedBoard.Name = board.Name
	}
	updatedBoard, err := b.BoardRepository.Update(selectedBoard)
	if err != nil {
		return nil, err
	}
	return updatedBoard, nil
}

func (b *BoardService) Delete(id *string) error {
	err := b.BoardRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BoardService) InviteUserToBoard(boardModel *model.InviteUserToBoardInput) (*[]entities.Board, error) {
	boardID := boardModel.BoardID.String()
	selectedBoard, err := b.BoardRepository.GetByID(&boardID)
	if err != nil {
		return nil, err
	}
	if selectedBoard == nil {
		return nil, errors.New("board not found")
	}
	if boardModel.UserID != selectedBoard.OwnerID {
		return nil, errors.New("you are not the owner of this board")
	}
	invitedBoard, err := b.BoardRepository.InviteUserToBoard(boardModel)
	if err != nil {
		return nil, err
	}
	if invitedBoard == nil {
		return nil, errors.New("failed to invite user to board")
	}
	return invitedBoard, nil
}

