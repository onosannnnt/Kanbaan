package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name string `gorm:"not null"`
	Description string `gorm:"not null"`
	ColumnID uuid.UUID `gorm:"not null" json:"column_id"`
	Column Column `gorm:"foreignKey:ColumnID;references:ID"`
	Assignee []User `gorm:"many2many:task_assignee;default:null"`
}