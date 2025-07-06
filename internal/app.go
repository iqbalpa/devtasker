package internal

import (
	"devtasker/internal/handler"
	"devtasker/internal/middleware"
	"devtasker/internal/repository"
	"devtasker/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func App(db *gorm.DB) *fiber.App {
	fmt.Println("Initializing App...")

	var tr repository.ITaskRepository = repository.New(db)
	var ts service.ITaskService = service.New(&tr)
	var th handler.TaskHandler = *handler.New(&ts)

	app := fiber.New()
	app.Use(middleware.Logger)

	app.Get("/doc/*", swagger.HandlerDefault)

	api := app.Group("/api")
	handler.TaskRouter(api, th)

	fmt.Println("App initiated successfully!")

	return app
}
