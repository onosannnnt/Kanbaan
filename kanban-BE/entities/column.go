package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Column struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name string `gorm:"not null"`
	BoardID uuid.UUID `gorm:"not null" json:"board_id"`
	Board Board `gorm:"foreignKey:BoardID;references:ID"`
}