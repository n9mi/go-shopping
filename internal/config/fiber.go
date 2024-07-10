package config

import (
	"fmt"
	"go-shopping/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(viperConfig *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      viperConfig.GetString("APP_NAME"),
		ErrorHandler: customErrorHandler(),
	})

	return app
}

func customErrorHandler() func(*fiber.Ctx, error) error {
	return func(c *fiber.Ctx, err error) error {
		response := new(model.MessageResponse)
		code := 500

		if err != nil {
			if errConv, ok := err.(validator.ValidationErrors); ok {
				for _, errItem := range errConv {
					switch errItem.Tag() {
					case "required":
						response.Messages = append(response.Messages,
							fmt.Sprintf("%s is required", errItem.Field()))
					case "min":
						response.Messages = append(response.Messages,
							fmt.Sprintf("%s is should be more than %s character", errItem.Field(), errItem.Param()))
					case "max":
						response.Messages = append(response.Messages,
							fmt.Sprintf("%s is should be less than %s character", errItem.Field(), errItem.Param()))
					case "email":
						response.Messages = append(response.Messages,
							fmt.Sprintf("%s should be a valid email", errItem.Field()))
					}
				}

				code = fiber.StatusBadRequest
			} else if errConv, ok := err.(*fiber.Error); ok {
				response.Messages = append(response.Messages, errConv.Message)

				code = errConv.Code
			}

			return c.Status(code).JSON(response)
		}

		return nil
	}
}
