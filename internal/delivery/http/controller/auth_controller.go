package controller

import (
	"go-shopping/internal/model"
	"go-shopping/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	AuthService service.AuthService
	Logger      *logrus.Logger
}

func NewAuthController(authService service.AuthService, logger *logrus.Logger) *AuthController {
	return &AuthController{
		AuthService: authService,
		Logger:      logger,
	}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Logger.Warnf("failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	if err := c.AuthService.Register(ctx.UserContext(), request); err != nil {
		return err
	}

	msgResponse := &model.MessageResponse{
		Messages: []string{"register succeed"},
	}
	return ctx.Status(fiber.StatusOK).JSON(msgResponse)
}
