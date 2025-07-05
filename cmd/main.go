package main

import (
	"devtasker/internal/handler"
	"devtasker/internal/middleware"
	"devtasker/internal/model"
	"devtasker/internal/repository"
	"devtasker/internal/service"
	"devtasker/internal/utils"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.InfoLogger.Println("App started!")

	db := utils.ConnectDb()
	db.AutoMigrate(&model.Task{})

	var tr repository.ITaskRepository = repository.New(db)
	var ts service.ITaskService = service.New(&tr)
	var th handler.TaskHandler = *handler.New(&ts)

	app := fiber.New()
	app.Use(middleware.Logger)
	api := app.Group("/api")
	handler.TaskRouter(api, th)

	app.Listen(":3000")
}
