package notificationUsecase

import "kanban/entities"

type NotificationUsecase interface {
	Create(notification *entities.Notification) (*entities.Notification, error)
	GetByID(id *string) (*entities.Notification, error)
	GetMyNotification(userID *string) (*[]entities.Notification, error)
	GetAll() (*[]entities.Notification, error)
	Update(notification *entities.Notification) (*entities.Notification, error)
	Delete(id *string) error
	GetUnreadCount(userID *string) (int64, error)
	MarkAsRead(notificationID *string) (*entities.Notification, error)
}