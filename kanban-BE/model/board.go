package model

import "github.com/google/uuid"

type InviteUserToBoardInput struct {
	BoardID uuid.UUID `json:"board_id"`
	UserID  uuid.UUID `json:"user_id"`
	Members []string  `json:"members"`
}
