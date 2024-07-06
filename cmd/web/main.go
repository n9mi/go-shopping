package main

import (
	"fmt"
	"go-shopping/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	viperCfg, err := config.NewViperConfig()
	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hello, world!")
	})

	app.Get("/check", func(c *fiber.Ctx) error {
		number := viperCfg.GetInt("NUMBER")

		return c.Status(fiber.StatusOK).SendString(
			fmt.Sprintf("%d x %d = %d", number, 10, number*10))
	})

	app.Listen(fmt.Sprintf(":%d", viperCfg.GetInt("APP_PORT")))
}
