package internal

import (
	"devtasker/internal/handler"
	"devtasker/internal/middleware"
	"devtasker/internal/repository"
	"devtasker/internal/service"
	"fmt"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func App(db *gorm.DB) *fiber.App {
	fmt.Println("Initializing App...")

	// Task
	var tr repository.ITaskRepository = repository.New(db)
	var ts service.ITaskService = service.New(tr)
	var th handler.TaskHandler = *handler.New(ts)

	// User & Auth
	var ur repository.IUserRepository = repository.NewUserRepo(db)
	var us service.IUserService = service.NewUserService(ur)
	var ah handler.AuthHandler = *handler.NewAuthHandler(us)

	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// Monitoring
	prometheus := fiberprometheus.New("devtasker")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Middleware
	app.Use(middleware.Logger)
	app.Use(middleware.Authorization)

	// API Doc
	app.Get("/doc/*", swagger.HandlerDefault)

	// API
	api := app.Group("/api")
	handler.TaskRouter(api, th)
	handler.AuthRouter(api, ah)

	fmt.Println("App initiated successfully!")

	return app
}
