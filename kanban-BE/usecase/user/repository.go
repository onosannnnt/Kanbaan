package UserUseCase

import "kanban/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	GetByID(id *string) (*entities.User, error)
	GetByEmail(email *string) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Delete(id *string) error
}
