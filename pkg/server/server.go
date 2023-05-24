package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/conf"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/handlers"
)

func NewServer(config conf.Config) *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("config", config)
		return c.Next()
	})

	app.Get("/healthcheck", handlers.Healthcheck)

	api := app.Group(config.EndpointsPefix)

	api.Post("/login/start", handlers.LoginStart)

	return app
}
