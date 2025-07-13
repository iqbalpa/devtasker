package handler

import (
	"devtasker/internal/dto"
	"devtasker/internal/service"
	"devtasker/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// ===== Router =====
func AuthRouter(api fiber.Router, ah AuthHandler) {
	api.Route("/auth", func(authRouter fiber.Router) {
		authRouter.Post("/register", ah.Register)
		authRouter.Post("/login", ah.Login)
	})
}

// ===== Handler =====
type AuthHandler struct {
	s service.IUserService
}

func NewAuthHandler(s service.IUserService) *AuthHandler {
	return &AuthHandler{
		s: s,
	}
}

// Register
// @Summary      Register a new user
// @Description  Register a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterUserRequest true "request body"
// @Success      200  {object}  model.User
// @Failure      500  {object}  error
// @Router       /api/auth/register [post]
func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	rur := new(dto.RegisterUserRequest)
	if err := c.BodyParser(rur); err != nil {
		utils.ErrorLogger.Println("Failed to parse the body:\n", c.Body())
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	u, err := ah.s.Register(rur.Name, rur.Username, rur.Passwrod)
	if err != nil {
		utils.ErrorLogger.Println("Failed to register the user:\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(u)
}

// Login
// @Summary      Login a user
// @Description  Login a user and get a JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginUserRequest true "request body"
// @Success      200  {object}  string
// @Failure      500  {object}  error
// @Router       /api/auth/login [post]
func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	lur := new(dto.LoginUserRequest)
	if err := c.BodyParser(lur); err != nil {
		utils.ErrorLogger.Println("Failed to parse the body:\n", c.Body())
		return c.Status(fiber.StatusInternalServerError).JSON(err)

	}
	token, err := ah.s.Login(lur.Username, lur.Passwrod)
	if err != nil {
		utils.ErrorLogger.Println("Failed to login the user:\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.JSON(token)
}
