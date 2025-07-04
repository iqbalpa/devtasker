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
