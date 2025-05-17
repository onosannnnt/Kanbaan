package notificationUsecase

import "kanban/entities"

type NotificationRepository interface {
	Create(notification *entities.Notification) (*entities.Notification, error)
	GetByID(id *string) (*entities.Notification, error)
	GetMyNotification(userID *string) (*[]entities.Notification, error)
	GetAll() (*[]entities.Notification, error)
	Update(notification *entities.Notification) (*entities.Notification, error)
	Delete(id *string) error
	GetUnreadCount(userID *string) (int64, error)
	MarkAsRead(notificationID *string) error
}

type NotificationService struct {
	NotificationRepository NotificationRepository
}

func NewNotificationUseCase(notificationRepository NotificationRepository) NotificationUsecase {
	return &NotificationService{
		NotificationRepository: notificationRepository,
	}
}

func (n *NotificationService) Create(notification *entities.Notification) (*entities.Notification, error) {
	createdNotification, err := n.NotificationRepository.Create(notification)
	if err != nil {
		return nil, err
	}
	return createdNotification, nil
}

func (n *NotificationService) GetByID(id *string) (*entities.Notification, error) {
	notification, err := n.NotificationRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (n *NotificationService) GetMyNotification(userID *string) (*[]entities.Notification, error) {
	notifications, err := n.NotificationRepository.GetMyNotification(userID)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (n *NotificationService) GetAll() (*[]entities.Notification, error) {
	notifications, err := n.NotificationRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

func (n *NotificationService) Update(notification *entities.Notification) (*entities.Notification, error) {
	updatedNotification, err := n.NotificationRepository.Update(notification)
	if err != nil {
		return nil, err
	}
	return updatedNotification, nil
}

func (n *NotificationService) Delete(id *string) error {
	err := n.NotificationRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotificationService) GetUnreadCount(userID *string) (int64, error) {
	count, err := n.NotificationRepository.GetUnreadCount(userID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (n *NotificationService) MarkAsRead(notificationID *string) (*entities.Notification, error) {
	err := n.NotificationRepository.MarkAsRead(notificationID)
	if err != nil {
		return nil, err
	}
	
	// Fetch the updated notification after marking it as read
	notification, err := n.NotificationRepository.GetByID(notificationID)
	if err != nil {
		return nil, err
	}
	
	return notification, nil
}
