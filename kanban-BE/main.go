package main

import (
	"fmt"
	"kanban/config"
	"kanban/entities"
	"kanban/router"
	"kanban/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println(config.DbSchema)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbSchema)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		panic("failed to create uuid-ossp extension")
	}
	entities.InitEntities(db)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	app.Get("/health", func(c fiber.Ctx) error {
		return utils.ResponseJSON(c, fiber.StatusOK, "Service is up and running", nil)
	})

	router.InitUserRouter(db, app)
	router.InitBoardRouter(db, app)
	router.InitColumnRouter(db, app)
	router.InitTaskRouter(db, app)
	router.InitNotificationRoute(db, app)

	app.Listen(":8000")
}
