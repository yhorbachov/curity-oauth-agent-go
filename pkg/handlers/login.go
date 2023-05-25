package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/yhorbachov/curity-oauth-agent-g/pkg/utils"
	"golang.org/x/exp/slices"
)

func LoginStart(c *fiber.Ctx) error {
	config := utils.GetConfig(c)

	if !slices.Contains(config.TrustedWebOrigins, c.Get("origin")) {
		return c.SendStatus(http.StatusForbidden)
	}

	var body struct {
		RedirectUri string `json:"redirectUri"`
	}

	if err := c.BodyParser(&body); err != nil || body.RedirectUri == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	return c.SendStatus(http.StatusOK)
}
