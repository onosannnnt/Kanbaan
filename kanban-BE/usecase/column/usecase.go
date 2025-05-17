package columnUsecase

import (
	"errors"
	"kanban/entities"
)

type ColumnUseCase interface {
	Create(column *entities.Column) (*entities.Column, error)
	GetByID(id *string) (*entities.Column, error)
	GetAll() (*[]entities.Column, error)
	Update(column *entities.Column) (*entities.Column, error)
	Delete(id *string) error
	GetByBoardID(boardID *string) (*[]entities.Column, error)
}

type ColumnService struct {
	ColumnRepository ColumnRepository
}

func NewColumnUseCase(columnRepository ColumnRepository) ColumnUseCase {
	return &ColumnService{
		ColumnRepository: columnRepository,
	}
}

func (c *ColumnService) Create(column *entities.Column) (*entities.Column, error) {
	createdColumn, err := c.ColumnRepository.Create(column)
	if err != nil {
		return nil, err
	}
	return createdColumn, nil
}

func (c *ColumnService) GetByID(id *string) (*entities.Column, error) {
	column, err := c.ColumnRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return column, nil
}

func (c *ColumnService) GetAll() (*[]entities.Column, error) {
	columns, err := c.ColumnRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return columns, nil
}

func (c *ColumnService) Update(column *entities.Column) (*entities.Column, error) {
	columnID := column.ID.String()
	selectedColumn, err := c.ColumnRepository.GetByID(&columnID)
	if err != nil {
		return nil, err
	}
	if selectedColumn == nil {
		return nil, errors.New("column not found")
	}
	if column.Name != "" {
		selectedColumn.Name = column.Name
	}
	updateColumn , err := c.ColumnRepository.Update(selectedColumn)
	if err != nil {
		return nil, err
	}
	return updateColumn, nil
}

func (c *ColumnService) Delete(id *string) error {
	err := c.ColumnRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *ColumnService) GetByBoardID(boardID *string) (*[]entities.Column, error) {
	columns, err := c.ColumnRepository.GetByBoardID(boardID)
	if err != nil {
		return nil, err
	}
	return columns, nil
}

