package UserUseCase

import (
	"kanban/config"
	"kanban/entities"
	BoardUseCase "kanban/usecase/board"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository  UserRepository
	BoardRepository BoardUseCase.BoardRepository
}

type UserUseCase interface {
	Create(user *entities.User) (*string, error)
	Login(email, password *string) (*string, error)
	GetUserByID(id *string) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(id *string) error
	Me(id *string) (*entities.User, error)
	GetMyBoards(id *string) (*[]entities.Board, error)
	GetColabBoards(id *string) (*[]entities.Board, error)
}

func NewUserUseCase(userRepository UserRepository, boardRepository BoardUseCase.BoardRepository) UserUseCase {
	return &UserService{
		UserRepository:  userRepository,
		BoardRepository: boardRepository,
	}
}

func (u *UserService) Create(user *entities.User) (*string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	createdUser, err := u.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	claim := jwt.MapClaims{
		"id":    createdUser.ID,
		"email": createdUser.Email,
		"exp":   time.Now().Add(time.Duration(config.ExpireTime) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (u *UserService) Login(email, password *string) (*string, error) {
	selectUser, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(selectUser.Password), []byte(*password)); err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"id":    selectUser.ID,
		"email": selectUser.Email,
		"exp":   time.Now().Add(time.Duration(config.ExpireTime) * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (u *UserService) GetUserByID(id *string) (*entities.User, error) {
	user, err := u.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) UpdateUser(user *entities.User) (*entities.User, error) {
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}
	id := user.ID.String()
	selectedUser, err := u.UserRepository.GetByID(&id)
	if err != nil {
		return nil, err
	}
	if user.Email != "" {
		selectedUser.Email = user.Email
	}
	if user.FirstName != "" {
		selectedUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		selectedUser.LastName = user.LastName
	}
	updatedUser, err := u.UserRepository.Update(selectedUser)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *UserService) DeleteUser(id *string) error {
	err := u.UserRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Me(id *string) (*entities.User, error) {
	user, err := u.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetMyBoards(id *string) (*[]entities.Board, error) {
	boards, err := u.BoardRepository.GetByOwnerID(id)
	if err != nil {
		return nil, err
	}
	return boards, nil
}

func (u *UserService) GetColabBoards(id *string) (*[]entities.Board, error) {
	boards, err := u.BoardRepository.GetColabBoards(id)
	if err != nil {
		return nil, err
	}
	return boards, nil
}
