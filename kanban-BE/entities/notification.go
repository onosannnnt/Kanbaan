package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title string `json:"title" gorm:"type:varchar(255)"`
	Message string `json:"message" gorm:"type:text"`
	Read bool `json:"read" gorm:"type:boolean;default:false"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User User `gorm:"foreignKey:UserID;references:ID"`
}