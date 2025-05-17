package BoardAdapter

import (
	"kanban/entities"
	"kanban/model"
	BoardUseCase "kanban/usecase/board"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type boardGorm struct {
	db *gorm.DB
}

func NewBoardGorm(db *gorm.DB) BoardUseCase.BoardRepository {
	return &boardGorm{
		db: db,
	}
}

func (b *boardGorm) Create(board *entities.Board) (*entities.Board, error) {
	if err := b.db.Create(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

func (b *boardGorm) GetByID(id *string) (*entities.Board, error) {
	var boards entities.Board
	if err := b.db.Preload("Members").Preload("Columns").Where("id = ?", id).First(&boards).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &boards, nil
}

func (b *boardGorm) GetAll() (*[]entities.Board, error) {
	var boards []entities.Board
	if err := b.db.Find(&boards).Error; err != nil {
		return nil, err
	}
	return &boards, nil
}

func (b *boardGorm) Update(board *entities.Board) (*entities.Board, error) {
	if err := b.db.Save(board).Error; err != nil {
		return nil, err
	}
	return board, nil
}

func (b *boardGorm) Delete(id *string) error {
	var board entities.Board
	if err := b.db.Where("id = ?", id).Delete(&board).Error; err != nil {
		return err
	}
	return nil
}

func (b *boardGorm) GetByOwnerID(ownerID *string) (*[]entities.Board, error) {
	var boards []entities.Board
	if err := b.db.Where("owner_id = ?", *ownerID).Find(&boards).Error; err != nil {
		return nil, err
	}
	return &boards, nil
}

func (b *boardGorm) InviteUserToBoard(boardModel *model.InviteUserToBoardInput) (*[]entities.Board, error) {
	var boards []entities.Board
	if err := b.db.Where("id = ?", boardModel.BoardID).Find(&boards).Error; err != nil {
		return nil, err
	}
	if len(boards) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	for _, memberEmail := range boardModel.Members {
		var user entities.User
		if err := b.db.Where("email = ?", memberEmail).First(&user).Error; err != nil {
			return nil, err
		}
		if user.ID == uuid.Nil {
			return nil, gorm.ErrRecordNotFound
		}
		// Assuming we're updating the first board found and it has a Members field
		boards[0].Members = append(boards[0].Members, user)
	}
	if err := b.db.Save(&boards[0]).Error; err != nil {
		return nil, err
	}

	return &boards, nil
}

func (b *boardGorm) GetColabBoards(userID *string) (*[]entities.Board, error) {
	var boards []entities.Board
	if err := b.db.Model(&entities.Board{}).
		Joins("JOIN board_members ON boards.id = board_members.board_id").
		Where("board_members.user_id = ?", userID).
		Find(&boards).Error; err != nil {
		return nil, err
	}
	if len(boards) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &boards, nil
}
