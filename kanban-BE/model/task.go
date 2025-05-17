package model

import "github.com/google/uuid"

type AssignTaskToUserInput struct {
	TaskID uuid.UUID `json:"task_id" validate:"required"`
	Assignee []string `json:"assignee_id" validate:"required"`
}