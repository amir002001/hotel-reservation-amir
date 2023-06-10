package api

import "github.com/gofiber/fiber/v2"

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON(map[string]any{"data": []string{"John", "Jacky"}})
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON(map[string]any{"data": "Jacky"})
}
