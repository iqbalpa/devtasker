package handler

import (
	"devtasker/internal/model"
	"devtasker/internal/service"

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

func (th *TaskHandler) CreateTask(c *fiber.Ctx) error {
	ctr := new(model.CreateTaskRequest)
	if err := c.BodyParser(ctr); err != nil {
		return err
	}
	t, err := th.s.CreateTask(ctr.Title, ctr.Description)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(t)
}

func (th *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := th.s.GetAllTasks()
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(tasks)
}

func (th *TaskHandler) GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := th.s.GetTaskByID(id)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(t)
}

func (th *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	b := new(model.UpdateTaskRequest)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	t, err := th.s.UpdateTask(
		id,
		b.Title,
		b.Description,
		b.Status,
	)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(t)
}

func (th *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := th.s.DeleteTask(id)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(t)
}
