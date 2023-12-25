package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/trenchesdeveloper/go-hotel/db"
	"github.com/trenchesdeveloper/go-hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	user, err := h.userStore.GetUserById(c.Context(), id)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return err
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "success",
		"users":   users,
	})
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if errors := params.Validate(); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	user, err := types.NewUserFromParams(params)

	if err != nil {
		return err
	}

	createdUser, err := h.userStore.CreateUser(c.Context(), user)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"user":    createdUser,
	})

}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	var (
		userId = c.Params("id")
		params types.UpdateUserParams
	)
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	paramsBson, errors := params.ToBson()

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}
	filter := bson.M{"_id": oid}

	if err := h.userStore.UpdateUser(c.Context(), filter, paramsBson);  err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("User with id %s updated successfully", userId),
	})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), userId); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
