package route

import (
	"go-shopping/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type RouteConfig struct {
	App             *fiber.App
	ControllerSetup *controller.ControllerSetup
}

func (c *RouteConfig) Setup() {
	route := c.App.Group("/api/v1.0")

	route.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	c.SetupAuthRoute(route)
}

func (c *RouteConfig) SetupAuthRoute(route fiber.Router) {
	auth := route.Group("/auth")

	auth.Post("/register", c.ControllerSetup.AuthController.Register)
}
