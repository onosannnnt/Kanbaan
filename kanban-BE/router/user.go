package router

import (
	BoardAdapter "kanban/adapter/board"
	UserAdepter "kanban/adapter/user"
	UserUseCase "kanban/usecase/user"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func InitUserRouter(db *gorm.DB, app *fiber.App) {
	userGorm := UserAdepter.NewUserGorm(db)
	boardGorm := BoardAdapter.NewBoardGorm(db)

	userService := UserUseCase.NewUserUseCase(userGorm, boardGorm)

	userHandler := UserAdepter.NewUserHandler(userService)

	user := app.Group("/users/")
	user.Post("/register", userHandler.Create)
	user.Post("/login", userHandler.Login)

	protected := user.Group("/")
	protected.Use(utils.IsAuth)

	protected.Post("/logout", userHandler.Logout)
	protected.Get("/me", userHandler.Me)
	protected.Get("/:id", userHandler.GetByID)
	protected.Put("/", userHandler.Update)
	protected.Delete("/", userHandler.Delete)
	protected.Get("/me/boards", userHandler.GetMyBoards)
	protected.Get("/me/boards/colab", userHandler.GetColabBoards)

}
