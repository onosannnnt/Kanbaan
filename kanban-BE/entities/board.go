package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name    string    `gorm:"not null"`
	OwnerID uuid.UUID `gorm:"not null"`
	Owner   User      `gorm:"foreignKey:OwnerID;references:ID"`
	Members []User    `gorm:"many2many:board_members;default:null"`
	Columns []Column  `gorm:"foreignKey:BoardID;references:ID"`
}
