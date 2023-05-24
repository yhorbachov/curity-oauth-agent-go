package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/utils"
	"golang.org/x/exp/slices"
)

func LoginStart(c *fiber.Ctx) error {
	config := utils.GetConfig(c)

	if !slices.Contains(config.TrustedWebOrigins, c.Get("Origin")) {
		return c.SendStatus(http.StatusForbidden)
	}

	return c.SendString("Login Start")
}
