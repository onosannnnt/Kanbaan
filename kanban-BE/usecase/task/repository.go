package taskUsecase

import (
	"kanban/entities"
	"kanban/model"
)

type TaskRepository interface {
	Create(task *entities.Task) (*entities.Task, error)
	GetByID(id *string) (*entities.Task, error)
	GetAll() (*[]entities.Task, error)
	Update(task *entities.Task) (*entities.Task, error)
	Delete(id *string) error
	GetByColumnID(columnID *string) (*[]entities.Task, error)
	AssignUser(taskModel *model.AssignTaskToUserInput) (*entities.Task, error) 
}