package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Healthcheck(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
