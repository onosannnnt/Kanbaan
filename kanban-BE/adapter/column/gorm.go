package columnAdapter

import (
	"kanban/entities"
	ColumnUsecase "kanban/usecase/column"

	"gorm.io/gorm"
)

type ColumnGorm struct {
	db *gorm.DB
}

func NewColumnGorm(db *gorm.DB) ColumnUsecase.ColumnRepository {
	return &ColumnGorm{
		db: db,
	}
}

func (c *ColumnGorm) Create(column *entities.Column) (*entities.Column, error) {
	if err := c.db.Create(column).Error; err != nil {
		return nil, err
	}
	return column, nil
}

func (c *ColumnGorm) GetByID(id *string) (*entities.Column, error) {
	var column entities.Column
	err := c.db.Where("id = ?", id).First(&column).Error
	if err != nil {
		return nil, err
	}
	return &column, nil
}
func (c *ColumnGorm) GetAll() (*[]entities.Column, error) {
	var columns []entities.Column
	err := c.db.Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return &columns, nil
}
func (c *ColumnGorm) Update(column *entities.Column) (*entities.Column, error) {
	err := c.db.Save(column).Error
	if err != nil {
		return nil, err
	}
	return column, nil
}
func (c *ColumnGorm) Delete(id *string) error {
	err := c.db.Delete(&entities.Column{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *ColumnGorm) GetByBoardID(boardID *string) (*[]entities.Column, error) {
	var columns []entities.Column
	err := c.db.Where("board_id = ?", boardID).Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return &columns, nil
}