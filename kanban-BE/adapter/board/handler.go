package BoardAdapter

import (
	"kanban/entities"
	"kanban/model"
	BoardUseCase "kanban/usecase/board"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type BoardHandler struct {
	boardUsecase BoardUseCase.BoardUseCase
}

func NewBoardHandler(boardUsecase BoardUseCase.BoardUseCase) *BoardHandler {
	return &BoardHandler{
		boardUsecase: boardUsecase,
	}
}

func (h *BoardHandler) CreateBoard(c fiber.Ctx) error {
	var board entities.Board
	if err := c.Bind().Body(&board); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	if board.Name == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Board name is required", nil)
	}
	uuidID, err := uuid.Parse(c.Locals("id").(string))
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	board.OwnerID = uuidID
	createdBoard, err := h.boardUsecase.Create(&board)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create board", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Board created successfully", createdBoard)
}

func (h *BoardHandler) GetBoardByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid board ID", nil)
	}

	board, err := h.boardUsecase.GetByID(&id)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Board not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get board", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Board retrieved successfully", board)
}

func (h *BoardHandler) GetAllBoards(c fiber.Ctx) error {
	boards, err := h.boardUsecase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get boards", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Boards retrieved successfully", boards)
}

func (h *BoardHandler) UpdateBoard(c fiber.Ctx) error {
	var board entities.Board
	var id = c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid board ID", nil)
	}
	if err := c.Bind().Body(&board); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	if board.Name == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Board name is required", nil)
	}
	uuidID, err := uuid.Parse(c.Locals("id").(string))
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	uuidBorardID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	board.ID = uuidBorardID
	board.OwnerID = uuidID
	updatedBoard, err := h.boardUsecase.Update(&board)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to update board", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Board updated successfully", updatedBoard)
}

func (h *BoardHandler) DeleteBoard(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid board ID", nil)
	}
	err := h.boardUsecase.Delete(&id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to delete board", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusNoContent, "Board deleted successfully", nil)
}

func (h *BoardHandler) InviteUserToBoard(c fiber.Ctx) error {
	id := c.Params("id")
	var boardModel model.InviteUserToBoardInput
	if err := c.Bind().Body(&boardModel); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid board ID", nil)
	}
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	boardModel.BoardID = uuidID
	userID := c.Locals("id").(string)
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	boardModel.UserID = uuidUserID
	invitedBoard, err := h.boardUsecase.InviteUserToBoard(&boardModel)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to invite user to board", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "User invited to board successfully", invitedBoard)
}
