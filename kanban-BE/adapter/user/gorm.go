package UserAdepter

import (
	"kanban/entities"
	UserUseCase "kanban/usecase/user"

	"gorm.io/gorm"
)

type userGorm struct {
	db *gorm.DB
}

func NewUserGorm(db *gorm.DB) UserUseCase.UserRepository {
	return &userGorm{
		db: db,
	}
}

func (u *userGorm) Create(user *entities.User) (*entities.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userGorm) GetByID(id *string) (*entities.User, error) {
	var user entities.User
	if err := u.db.Omit("Password").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userGorm) GetByEmail(email *string) (*entities.User, error) {
	var user entities.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userGorm) Update(user *entities.User) (*entities.User, error) {
	if err := u.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userGorm) Delete(id *string) error {
	var user entities.User
	if err := u.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
