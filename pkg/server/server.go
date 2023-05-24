package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/conf"
)

func NewServer(config conf.Config) *fiber.App {
	app := fiber.New()

	api := app.Group(config.EndpointsPefix)

	api.Get(config.EndpointsPefix+"/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	return app
}
