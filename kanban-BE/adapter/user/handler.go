package UserAdepter

import (
	"kanban/entities"
	UserUseCase "kanban/usecase/user"
	"kanban/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUseCase UserUseCase.UserUseCase
}

func NewUserHandler(userUseCase UserUseCase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) Create(c fiber.Ctx) error {
	var user entities.User
	if err := c.Bind().Body(&user); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	user.Email = strings.ToLower(user.Email)
	if user.Email == "" || user.Password == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Email and password are required", nil)
	}
	createdUser, err := h.userUseCase.Create(&user)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    *createdUser,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	}
	c.Cookie(&cookie)
	return utils.ResponseJSON(c, fiber.StatusCreated, "User created successfully", createdUser)
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}

	user, err := h.userUseCase.GetUserByID(&id)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "User not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get user", err)
	}

	return utils.ResponseJSON(c, fiber.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	var user entities.User
	if err := c.Bind().Body(&user); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	id := c.Locals("id").(string)
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	user.ID = uuidID
	updatedUser, err := h.userUseCase.UpdateUser(&user)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return utils.ResponseJSON(c, fiber.StatusOK, "User updated successfully", updatedUser)
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
	id := c.Locals("id").(string)
	err := h.userUseCase.DeleteUser(&id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to delete user", err.Error())
	}
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Expires: time.Now().Add(-time.Hour * 24),
		Value:   "",
	})
	return utils.ResponseJSON(c, fiber.StatusNoContent, "User deleted successfully", nil)
}
func (h *UserHandler) Login(c fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var loginRequest LoginRequest
	if err := c.Bind().Body(&loginRequest); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request", err)
	}
	if loginRequest.Email == "" || loginRequest.Password == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Email and password are required", nil)
	}
	loginRequest.Email = strings.ToLower(loginRequest.Email)
	token, err := h.userUseCase.Login(&loginRequest.Email, &loginRequest.Password)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusUnauthorized, "Invalid email or password", err.Error())
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    *token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	}
	c.Cookie(&cookie)
	return utils.ResponseJSON(c, fiber.StatusOK, "Login successful", fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) Logout(c fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Expires: time.Now().Add(-time.Hour * 24),
		Value:   "",
	})
	return utils.ResponseJSON(c, fiber.StatusOK, "Logout successful", nil)
}

func (h *UserHandler) Me(c fiber.Ctx) error {
	id := c.Locals("id").(string)
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}

	user, err := h.userUseCase.GetUserByID(&id)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ResponseJSON(c, fiber.StatusNotFound, "User not found", nil)
		}
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get user", err)
	}

	return utils.ResponseJSON(c, fiber.StatusOK, "User retrieved successfully", user)
}

func (h *UserHandler) GetMyBoards(c fiber.Ctx) error {
	id := c.Locals("id").(string)
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	boards, err := h.userUseCase.GetMyBoards(&id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get boards", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Boards retrieved successfully", boards)
}

func (h *UserHandler) GetColabBoards(c fiber.Ctx) error {
	id := c.Locals("id").(string)
	if id == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid user ID", nil)
	}
	boards, err := h.userUseCase.GetColabBoards(&id)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get boards", err.Error())
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Boards retrieved successfully", boards)
}
