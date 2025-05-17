package columnAdapter

import (
	"fmt"
	"kanban/entities"
	ColumnUsecase "kanban/usecase/column"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ColumnHandler struct {
	ColumnUsecase ColumnUsecase.ColumnRepository
}

func NewColumnHandler(columnUsecase ColumnUsecase.ColumnRepository) *ColumnHandler {
	return &ColumnHandler{
		ColumnUsecase: columnUsecase,
	}
}

func (h *ColumnHandler) Create(c fiber.Ctx) error {
	var column entities.Column
	if err := c.Bind().Body(&column); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	fmt.Println("column", column.BoardID)

	newColumn, err := h.ColumnUsecase.Create(&column)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Column not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create column", err)
	}

	return utils.ResponseJSON(c, fiber.StatusCreated, "Column created successfully", newColumn)
}

func (h *ColumnHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid column ID", nil)
	}

	column, err := h.ColumnUsecase.GetByID(&id)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Column not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get column", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Column retrieved successfully", column)
}

func (h *ColumnHandler) GetAll(c fiber.Ctx) error {
	columns, err := h.ColumnUsecase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get columns", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Columns retrieved successfully", columns)
}

func (h *ColumnHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid column ID", nil)
	}

	var column entities.Column
	if err := c.Bind().Body(&column); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	uuidColumnID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid column ID", nil)
	}
	column.ID = uuidColumnID

	updatedColumn, err := h.ColumnUsecase.Update(&column)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Column not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to update column", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Column updated successfully", updatedColumn)
}

func (h *ColumnHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid column ID", nil)
	}

	err := h.ColumnUsecase.Delete(&id)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Column not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to delete column", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Column deleted successfully", nil)
}

func (h *ColumnHandler) GetByBoardID(c fiber.Ctx) error {
	boardID := c.Params("board_id")
	if boardID == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid board ID", nil)
	}

	columns, err := h.ColumnUsecase.GetByBoardID(&boardID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Columns not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get columns", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Columns retrieved successfully", columns)
}