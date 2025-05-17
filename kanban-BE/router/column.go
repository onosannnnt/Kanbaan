package router

import (
	columnAdapter "kanban/adapter/column"
	columnUsecase "kanban/usecase/column"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitColumnRouter(db *gorm.DB, app *fiber.App) {
	ColumnGorm := columnAdapter.NewColumnGorm(db)
	ColumnUseCase := columnUsecase.NewColumnUseCase(ColumnGorm)
	ColumnHandler := columnAdapter.NewColumnHandler(ColumnUseCase)
	
	column := app.Group("/columns")
	
	protected := column.Group("/")
	protected.Use(utils.IsAuth)

	protected.Post("/", ColumnHandler.Create)
	protected.Get("/:id", ColumnHandler.GetByID)
	protected.Get("/", ColumnHandler.GetAll)
	protected.Put("/:id", ColumnHandler.Update)
	protected.Delete("/:id", ColumnHandler.Delete)
	protected.Get("/board/:board_id", ColumnHandler.GetByBoardID)
}