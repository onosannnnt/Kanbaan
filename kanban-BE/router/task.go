package router

import (
	notificationAdapter "kanban/adapter/notification"
	taskAdapter "kanban/adapter/task"
	UserAdepter "kanban/adapter/user"
	taskUsecase "kanban/usecase/task"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitTaskRouter(db *gorm.DB, app *fiber.App) {
	taskRepository := taskAdapter.NewTaskGorm(db)
	userGorm := UserAdepter.NewUserGorm(db)
	notificationGorm := notificationAdapter.NewNotificationGorm(db)

	taskUseCase := taskUsecase.NewTaskUseCase(taskRepository, userGorm, notificationGorm)
	taskHandler := taskAdapter.NewTaskHandler(taskUseCase)

	task := app.Group("/tasks")

	protected := task.Group("/")
	protected.Use(utils.IsAuth)

	protected.Post("/", taskHandler.Create)
	protected.Get("/:id", taskHandler.GetByID)
	protected.Get("/", taskHandler.GetAll)
	protected.Put("/:id", taskHandler.Update)
	protected.Delete("/:id", taskHandler.Delete)
	protected.Get("/column/:columnID", taskHandler.GetByColumnID)
	protected.Get("/column", taskHandler.GetByColumnID)
	protected.Post("/:id/assigns", taskHandler.AssignUser)
	
}