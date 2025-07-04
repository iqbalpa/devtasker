package main

import (
	"devtasker/internal/handler"
	"devtasker/internal/repository"
	"devtasker/internal/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Welcome!")
	var tr repository.ITaskRepository = repository.New()
	var ts service.ITaskService = service.New(&tr)
	var th handler.TaskHandler = *handler.New(&ts)

	app := fiber.New()
	api := app.Group("/api")
	handler.TaskRouter(api, th)

	app.Listen(":3000")
}
