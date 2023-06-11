package api

import (
	"hotel-amir/db"
	"hotel-amir/types"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userStore db.UserStore
}

const bcryptCost = 12

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	createUserParams := types.CreateUserParams{}
	err := c.BodyParser(&createUserParams)
	if err != nil {
		return err
	}
	// TODO SALT?
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(createUserParams.Password),
		bcryptCost,
	)
	if err != nil {
		return err
	}
	user := types.User{
		FirstName:         createUserParams.FirstName,
		LastName:          createUserParams.LastName,
		Email:             createUserParams.Email,
		EncryptedPassword: string(encryptedPassword),
	}
	// idempotent
	// status codes TODO
	insertedUser, err := h.userStore.CreateUser(c.Context(), &user)
	if err != nil {
		return err
	}
	return c.JSON(map[string]any{"data": insertedUser})
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetAllUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(map[string]any{"data": users})
}

func (h *UserHandler) HandleGetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(map[string]any{"data": user})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	deletedCount, err := h.userStore.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(map[string]any{"data": deletedCount})
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userParams := types.UpdateUserParams{}
	err := c.BodyParser(&userParams)
	desired := types.User{
		FirstName: userParams.FirstName,
		LastName:  userParams.LastName,
		Email:     userParams.Email,
	}
	if err != nil {
		return err
	}
	updatedId, err := h.userStore.UpdateUser(c.Context(), id, &desired)
	if err != nil {
		return err
	}
	return c.JSON(map[string]any{"data": updatedId})
}
