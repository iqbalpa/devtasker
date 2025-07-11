package handler

import (
	"devtasker/internal/model"
	"devtasker/internal/service"
	"devtasker/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// ===== Router =====
func TaskRouter(api fiber.Router, th TaskHandler) {
	api.Route("/task", func(taskRouter fiber.Router) {
		taskRouter.Post("/", th.CreateTask)
		taskRouter.Get("/", th.GetAllTasks)
		taskRouter.Get("/:id", th.GetTaskByID)
		taskRouter.Patch("/:id", th.UpdateTask)
		taskRouter.Delete("/:id", th.DeleteTask)
	})
}

// ===== Handler =====
type TaskHandler struct {
	s service.ITaskService
}

func New(s *service.ITaskService) *TaskHandler {
	return &TaskHandler{
		s: *s,
	}
}

// CreateTask
// @Summary      Create a new task
// @Description  Create a new task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        request body model.CreateTaskRequest true "request body"
// @Success      200  {object}  model.Task
// @Failure      500  {object}  error
// @Router       /api/task [post]
func (th *TaskHandler) CreateTask(c *fiber.Ctx) error {
	ctr := new(model.CreateTaskRequest)
	if err := c.BodyParser(ctr); err != nil {
		utils.ErrorLogger.Println("Failed to parse the body:\n", c.Body())
		return err
	}
	t, err := th.s.CreateTask(ctr.Title, ctr.Description)
	if err != nil {
		utils.ErrorLogger.Println("Failed to create a new task:\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(t)
}

// GetAllTasks
// @Summary      Get all tasks
// @Description  Retrieve a list of all tasks
// @Tags         task
// @Produce      json
// @Success      200  {array}  model.Task
// @Failure      500  {object}  error
// @Router       /api/task [get]
func (th *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := th.s.GetAllTasks()
	if err != nil {
		utils.ErrorLogger.Println("Failed to get all tasks:\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(tasks)
}

// GetTaskByID
// @Summary      Get task by ID
// @Description  Retrieve a task by its ID
// @Tags         task
// @Produce      json
// @Param        id   path      string  true  "Task ID"
// @Success      200  {object}  model.Task
// @Failure      404  {object}  error
// @Router       /api/task/{id} [get]
func (th *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := th.s.GetTaskByID(id)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to get task with id %s:\n%s", id, err)
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.JSON(t)
}

// UpdateTask
// @Summary      Update a task
// @Description  Update the title, description, or status of a task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id      path      string                  true  "Task ID"
// @Param        request body      model.UpdateTaskRequest true  "Update Task Body"
// @Success      200     {object}  model.Task
// @Failure      500     {object}  error
// @Router       /api/task/{id} [patch]
func (th *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	b := new(model.UpdateTaskRequest)
	if err := c.BodyParser(b); err != nil {
		utils.ErrorLogger.Println("Failed to parse the body:\n", c.Body())
		return err
	}
	t, err := th.s.UpdateTask(
		id,
		b.Title,
		b.Description,
		b.Status,
	)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to update task with id %s:\n%s", id, err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(t)
}

// DeleteTask
// @Summary      Delete a task
// @Description  Delete a task by its ID
// @Tags         task
// @Produce      json
// @Param        id   path      string  true  "Task ID"
// @Success      200  {object}  model.Task
// @Failure      500  {object}  error
// @Router       /api/task/{id} [delete]
func (th *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := th.s.DeleteTask(id)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to delete task with id %s:\n%s", id, err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(t)
}
