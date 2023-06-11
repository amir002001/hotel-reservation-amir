package api

import (
	"hotel-amir/db"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(map[string]any{"data": users})
}

func (h *UserHandler) HandleGetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserById(id)
	if err != nil {
		return err
	}
	return c.JSON(map[string]any{"data": user})
}
