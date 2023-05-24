package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/conf"
)

func GetConfig(c *fiber.Ctx) conf.Config {
	return c.Locals("config").(conf.Config)
}
