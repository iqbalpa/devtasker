package middleware

import (
	"devtasker/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)
	utils.InfoLogger.Printf(
		"%s %s | %d | %s | %s",
		c.Method(),
		c.Path(),
		c.Response().StatusCode(),
		duration,
		c.IP(),
	)
	return err
}
