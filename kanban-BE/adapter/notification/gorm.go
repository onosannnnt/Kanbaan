package notificationAdapter

import (
	"kanban/entities"
	nottificationUsecase "kanban/usecase/notifiaction"

	"gorm.io/gorm"
)


type NotificationGorm struct {
	DB *gorm.DB
}

func NewNotificationGorm(db *gorm.DB) nottificationUsecase.NotificationRepository {
	return &NotificationGorm{
		DB: db,
	}
}

func (n *NotificationGorm) Create(notification *entities.Notification) (*entities.Notification, error) {
	if err := n.DB.Create(notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (n *NotificationGorm) GetByID(id *string) (*entities.Notification, error) {
	var notification entities.Notification
	err := n.DB.Where("id = ?", id).First(&notification).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

func (n *NotificationGorm) GetMyNotification(userID *string) (*[]entities.Notification, error) {
	var notifications []entities.Notification
	err := n.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return &notifications, nil
}

func (n *NotificationGorm) GetAll() (*[]entities.Notification, error) {
	var notifications []entities.Notification
	err := n.DB.Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return &notifications, nil
}

func (n *NotificationGorm) Update(notification *entities.Notification) (*entities.Notification, error) {
	err := n.DB.Save(notification).Error
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (n *NotificationGorm) Delete(id *string) error {
	err := n.DB.Delete(&entities.Notification{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (n *NotificationGorm) GetUnreadCount(userID *string) (int64, error) {
	var count int64
	err := n.DB.Model(&entities.Notification{}).Where("user_id = ? AND read = ?", userID, false).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (n *NotificationGorm) MarkAsRead(notificationID *string) error {
	err := n.DB.Model(&entities.Notification{}).Where("id = ?", notificationID).Update("read", true).Error
	if err != nil {
		return err
	}
	return nil
}