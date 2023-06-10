package api

import (
	"hotel-amir/types"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	users := []types.User{
		{FirstName: "Amir", LastName: "Afshari"},
		{FirstName: "Jacky", LastName: "Johanson"},
	}
	return c.JSON(map[string]any{"data": users})
}

func HandleGetUserById(c *fiber.Ctx) error {
	user := types.User{FirstName: "Amir", LastName: "Afshari"}
	return c.JSON(map[string]any{"data": user})
}
