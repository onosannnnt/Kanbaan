package router

import (
	BoardAdapter "kanban/adapter/board"
	BoardUseCase "kanban/usecase/board"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitBoardRouter(db *gorm.DB, app *fiber.App) {
	boardGorm := BoardAdapter.NewBoardGorm(db)

	boardUseCase := BoardUseCase.NewBoardUseCase(boardGorm)
	boardHandler := BoardAdapter.NewBoardHandler(boardUseCase)

	board := app.Group("/boards")

	protected := board.Group("/")
	protected.Use(utils.IsAuth)

	protected.Post("/", boardHandler.CreateBoard)
	protected.Get("/:id", boardHandler.GetBoardByID)
	protected.Get("/", boardHandler.GetAllBoards)
	protected.Put("/:id", boardHandler.UpdateBoard)
	protected.Delete("/:id", boardHandler.DeleteBoard)
	protected.Post("/:id/invites", boardHandler.InviteUserToBoard)
}
