package middleware

import (
	"context"
	"devtasker/internal/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Authorization(c *fiber.Ctx) error {
	if strings.Contains(c.Path(), "auth") ||
		strings.Contains(c.Path(), "doc") ||
		strings.Contains(c.Path(), "health") {
		return c.Next()
	}

	bearerToken := c.Get("Authorization")
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or malformed token",
		})
	}

	token := strings.Split(bearerToken, " ")[1]
	claims, ok := utils.ExtractClaims(token)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token invalid",
		})
	}

	if expRaw, ok := claims["exp"].(float64); ok {
		expTime := time.Unix(int64(expRaw), 0)
		if time.Now().After(expTime) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token expired",
			})
		}
	}

	ctx := context.WithValue(c.Context(), utils.UsernameKey, claims["username"])
	ctx = context.WithValue(ctx, utils.NameKey, claims["name"])
	c.SetUserContext(ctx)

	return c.Next()
}
