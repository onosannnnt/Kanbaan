package columnUsecase

import "kanban/entities"

type ColumnRepository interface {
	Create(column *entities.Column) (*entities.Column, error)
	GetByID(id *string) (*entities.Column, error)
	GetAll() (*[]entities.Column, error)
	Update(column *entities.Column) (*entities.Column, error)
	Delete(id *string) error
	GetByBoardID(boardID *string) (*[]entities.Column, error)
}