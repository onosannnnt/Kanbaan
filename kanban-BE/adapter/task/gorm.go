package taskAdapter

import (
	"fmt"
	"kanban/entities"
	"kanban/model"
	taskUsecase "kanban/usecase/task"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type TaskGorm struct {
	DB *gorm.DB
}

func NewTaskGorm(db *gorm.DB) taskUsecase.TaskRepository {
	return &TaskGorm{
		DB: db,
	}
}

func (t *TaskGorm) Create(task *entities.Task) (*entities.Task, error) {
	if err := t.DB.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskGorm) GetByID(id *string) (*entities.Task, error) {
	var task entities.Task
	err := t.DB.Preload("Assignee").Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *TaskGorm) GetAll() (*[]entities.Task, error) {
	var tasks []entities.Task
	err := t.DB.Preload("Assignee").Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (t *TaskGorm) Update(task *entities.Task) (*entities.Task, error) {
	err := t.DB.Save(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskGorm) Delete(id *string) error {
	err := t.DB.Delete(&entities.Task{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *TaskGorm) GetByColumnID(columnID *string) (*[]entities.Task, error) {
	var tasks []entities.Task
	err := t.DB.Where("column_id = ?", columnID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (t *TaskGorm) AssignUser(taskModel *model.AssignTaskToUserInput) (*entities.Task, error) {
	var task entities.Task
	err := t.DB.Where("id = ?", taskModel.TaskID).First(&task).Error
	if err != nil {
		return nil, err
	}
	if task.ID == uuid.Nil {
		return nil, gorm.ErrRecordNotFound
	}
	fmt.Println(taskModel.Assignee)
	for _, memberEmail := range taskModel.Assignee {
		var user entities.User
		if err := t.DB.Where("email = ?", memberEmail).First(&user).Error; err != nil {
			return nil, err
		}
		if user.ID == uuid.Nil {
			return nil, gorm.ErrRecordNotFound
		}
		task.Assignee = append(task.Assignee, user)
	}
	if err := t.DB.Save(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}