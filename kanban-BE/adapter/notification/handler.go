package notificationAdapter

import (
	"kanban/entities"
	notificationUsecase "kanban/usecase/notifiaction"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type notificationAdapter struct {
	notificationUsecase notificationUsecase.NotificationUsecase
}

func NewNotificationAdapter(notificationUsecase notificationUsecase.NotificationUsecase) *notificationAdapter {
	return &notificationAdapter{
		notificationUsecase: notificationUsecase,
	}
}

func (n *notificationAdapter) Create(c fiber.Ctx) error {
	var notification entities.Notification
	if err := c.Bind().Body(&notification); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	createdNotification, err := n.notificationUsecase.Create(&notification)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create notification", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusCreated, "Notification created successfully", createdNotification)
}

func (n *notificationAdapter) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid notification ID", nil)
	}

	notification, err := n.notificationUsecase.GetByID(&id)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Notification not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get notification", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Notification retrieved successfully", notification)
}

func (n *notificationAdapter) GetMyNotification(c fiber.Ctx) error {
	userID := c.Locals("id")
	if userID == nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "User ID is required", nil)
	}
	strUserID := userID.(string)
	notifications, err := n.notificationUsecase.GetMyNotification(&strUserID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get notifications", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Notifications retrieved successfully", notifications)
}

func (n *notificationAdapter) GetAll(c fiber.Ctx) error {
	notifications, err := n.notificationUsecase.GetAll()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get notifications", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Notifications retrieved successfully", notifications)
}

func (n *notificationAdapter) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid notification ID", nil)
	}

	var notification entities.Notification
	if err := c.Bind().Body(&notification); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid notification ID", nil)
	}
	notification.ID = uuidID
	updatedNotification, err := n.notificationUsecase.Update(&notification)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to update notification", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Notification updated successfully", updatedNotification)
}

func (n *notificationAdapter) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid notification ID", nil)
	}

	err := n.notificationUsecase.Delete(&id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to delete notification", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Notification deleted successfully", nil)
}

func (n *notificationAdapter) GetUnreadCount(c fiber.Ctx) error {
	userID := c.Locals("id")
	if userID == nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "User ID is required", nil)
	}
	strUserID := userID.(string)
	count, err := n.notificationUsecase.GetUnreadCount(&strUserID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get unread count", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Unread count retrieved successfully", count)
}

func (n *notificationAdapter) MarkAsRead(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid notification ID", nil)
	}

	notification, err := n.notificationUsecase.MarkAsRead(&id)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "Notification not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to mark notification as read", err)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Notification marked as read successfully", notification)
}
