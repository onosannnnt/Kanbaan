package router

import (
	notificationAdapter "kanban/adapter/notification"
	notificationUsecase "kanban/usecase/notifiaction"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitNotificationRoute(db *gorm.DB, app *fiber.App) {
	notificationGorm := notificationAdapter.NewNotificationGorm(db)
	notificationService := notificationUsecase.NewNotificationUseCase(notificationGorm)

	notificationHandler := notificationAdapter.NewNotificationAdapter(notificationService)

	notification := app.Group("/notifications")
	protected := notification.Group("/")
	protected.Use(utils.IsAuth)
	
	protected.Get("/", notificationHandler.GetAll)
	protected.Get("/my", notificationHandler.GetMyNotification)
	protected.Get("/:id", notificationHandler.GetByID)
	protected.Post("/", notificationHandler.Create)
	protected.Put("/:id", notificationHandler.Update)
	protected.Delete("/:id", notificationHandler.Delete)

	protected.Get("/unread/count", notificationHandler.GetUnreadCount)
	protected.Put("/mark-as-read/:id", notificationHandler.MarkAsRead)
	
}