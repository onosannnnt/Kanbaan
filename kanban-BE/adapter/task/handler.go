package taskAdapter

import (
	"kanban/entities"
	"kanban/model"
	taskUsecase "kanban/usecase/task"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)


type TaskHandler struct {
	TaskUseCase taskUsecase.TaskUseCase
}

func NewTaskHandler(taskUseCase taskUsecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		TaskUseCase: taskUseCase,
	}
}

func (h *TaskHandler) Create(c fiber.Ctx) error {
	var task entities.Task
	if err := c.Bind().Body(&task); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	if task.Name == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Task name is required", nil)
	}
	createdTask, err := h.TaskUseCase.Create(&task)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create task", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Task created successfully", createdTask)
}

func (h *TaskHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid task ID", nil)
	}

	task, err := h.TaskUseCase.GetByID(&id)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Task not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get task", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Task retrieved successfully", task)
}

func (h *TaskHandler) GetAll(c fiber.Ctx) error {
	tasks, err := h.TaskUseCase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get tasks", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Tasks retrieved successfully", tasks)
}

func (h *TaskHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid task ID", nil)
	}

	var task entities.Task
	if err := c.Bind().Body(&task); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	task.ID = uuidID

	updatedTask, err := h.TaskUseCase.Update(&task)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to update task", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Task updated successfully", updatedTask)
}

func (h *TaskHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid task ID", nil)
	}

	err := h.TaskUseCase.Delete(&id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to delete task", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Task deleted successfully", nil)
}

func (h *TaskHandler) GetByColumnID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid column ID", nil)
	}

	tasks, err := h.TaskUseCase.GetByColumnID(&id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get tasks", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Tasks retrieved successfully", tasks)
}

func (h *TaskHandler) AssignUser(c fiber.Ctx) error {
	var id = c.Params("id")
	var taskModel model.AssignTaskToUserInput
	if err := c.Bind().Body(&taskModel); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	if id == "" || len(taskModel.Assignee) == 0 {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Task ID and User ID are required", nil)
	}
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid task ID", nil)
	}
	taskModel.TaskID = uuidID
	assignedTask, err := h.TaskUseCase.AssingTaskToUser(&taskModel)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to assign user to task", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "User assigned to task successfully", assignedTask)
}