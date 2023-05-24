package handlers

import "github.com/gofiber/fiber/v2"

func LoginStart(c *fiber.Ctx) error {
	return c.SendString("Login Start")
}
