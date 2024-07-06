package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// viperCfg, err := config.NewViperConfig()
	// if err != nil {
	// 	panic(err)
	// }

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hello, world!")
	})

	app.Get("/check", func(c *fiber.Ctx) error {
		numberStr := os.Getenv("NUMBER")
		number, _ := strconv.Atoi(numberStr)

		return c.Status(fiber.StatusOK).SendString(
			fmt.Sprintf("%d x %d = %d", number, 10, number*10))
	})

	app.Listen(":3000")
}
