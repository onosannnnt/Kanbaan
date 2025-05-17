package taskUsecase

import (
	"fmt"
	"kanban/entities"
	"kanban/model"
	notificationUsecase "kanban/usecase/notifiaction"
	UserUseCase "kanban/usecase/user"

	"github.com/google/uuid"
)

type TaskUseCase interface {
	Create(task *entities.Task) (*entities.Task, error)
	GetByID(id *string) (*entities.Task, error)
	GetAll() (*[]entities.Task, error)
	Update(task *entities.Task) (*entities.Task, error)
	Delete(id *string) error
	GetByColumnID(columnID *string) (*[]entities.Task, error)
	AssingTaskToUser(taskModel *model.AssignTaskToUserInput) (*entities.Task, error)
}

type TaskService struct {
	TaskRepository TaskRepository
	UserRepository UserUseCase.UserRepository
	NotificationUsecase notificationUsecase.NotificationRepository
}

func NewTaskUseCase(taskRepository TaskRepository, userRepository UserUseCase.UserRepository, NotificationUsecase notificationUsecase.NotificationRepository) TaskUseCase {
	return &TaskService{
		TaskRepository: taskRepository,
		UserRepository: userRepository,
		NotificationUsecase: NotificationUsecase,
	}
}

func (t *TaskService) Create(task *entities.Task) (*entities.Task, error) {
	createdTask, err := t.TaskRepository.Create(task)
	if err != nil {
		return nil, err
	}
	return createdTask, nil
}

func (t *TaskService) GetByID(id *string) (*entities.Task, error) {
	task, err := t.TaskRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskService) GetAll() (*[]entities.Task, error) {
	tasks, err := t.TaskRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskService) Update(task *entities.Task) (*entities.Task, error) {
	taskID := task.ID.String()
	selectedTask, err := t.TaskRepository.GetByID(&taskID)
	if err != nil {
		return nil, err
	}
	if selectedTask == nil {
		return nil, err
	}
	if task.Name != "" {
		selectedTask.Name = task.Name
	}
	if task.Description != "" {
		selectedTask.Description = task.Description
	}
	if task.ColumnID != uuid.Nil {
		selectedTask.ColumnID = task.ColumnID
	}
	updateTask, err := t.TaskRepository.Update(selectedTask)
	if err != nil {
		return nil, err
	}
	return updateTask, nil
}

func (t *TaskService) Delete(id *string) error {
	err := t.TaskRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskService) GetByColumnID(columnID *string) (*[]entities.Task, error) {
	tasks, err := t.TaskRepository.GetByColumnID(columnID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskService) AssingTaskToUser(taskModel *model.AssignTaskToUserInput) (*entities.Task, error) {
	taskID := taskModel.TaskID.String()
	task, err := t.TaskRepository.GetByID(&taskID)
	if err != nil {
		return nil, err
	}
	if task == nil {
		return nil, err
	}
	var notification entities.Notification
	notification.ID = uuid.New()
	selectedUser,err := t.UserRepository.GetByEmail(&taskModel.Assignee[0])
	if err != nil {
		return nil, err
	}
	if selectedUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	assignedUser := selectedUser.ID
	notification.UserID = assignedUser
	notification.Title = "Task Assigned"
	notification.Message = "You have been assigned a task"
	_, err = t.NotificationUsecase.Create(&notification)
	if err != nil {
		return nil, err
	}
	assignTask, err := t.TaskRepository.AssignUser(taskModel)
	if err != nil {
		return nil, err
	}
	if assignTask == nil {
		return nil, err
	}
	return assignTask, nil
	
}