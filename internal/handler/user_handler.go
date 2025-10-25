package handler

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"user-api/internal/models"
	"user-api/internal/service"
)

type UserHandler struct {
	userService service.UserService
	validator   *validator.Validate
	logger      *zap.Logger
}

func NewUserHandler(userService service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator.New(),
		logger:      logger,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("bad request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrInvalidInput)
	}

	if err := h.validator.Struct(&req); err != nil {
		h.logger.Error("validation error", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrValidationFailed)
	}

	user, err := h.userService.CreateUser(c.Context(), &req)
	if err != nil {
		h.logger.Error("couldn't create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrInternalServer)
	}

	h.logger.Info("user created", zap.Int32("user_id", user.ID))
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.logger.Error("invalid user id", zap.String("id", idStr), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrInvalidInput)
	}

	user, err := h.userService.GetUser(c.Context(), int32(id))
	if err != nil {
		if err == models.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(err)
		}
		h.logger.Error("failed to get user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrInternalServer)
	}

	h.logger.Info("user retrieved", zap.Int32("user_id", user.ID))
	return c.JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userService.ListUsers(c.Context())
	if err != nil {
		h.logger.Error("couldn't list users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrInternalServer)
	}

	h.logger.Info("users listed", zap.Int("count", len(users)))
	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.logger.Error("invalid user id", zap.String("id", idStr), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrInvalidInput)
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("bad request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrInvalidInput)
	}

	if err := h.validator.Struct(&req); err != nil {
		h.logger.Error("validation error", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrValidationFailed)
	}

	user, err := h.userService.UpdateUser(c.Context(), int32(id), &req)
	if err != nil {
		if err == models.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(err)
		}
		h.logger.Error("couldn't update user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrInternalServer)
	}

	h.logger.Info("user updated", zap.Int32("user_id", user.ID))
	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		h.logger.Error("invalid user id", zap.String("id", idStr), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrInvalidInput)
	}

	err = h.userService.DeleteUser(c.Context(), int32(id))
	if err != nil {
		if err == models.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(err)
		}
		h.logger.Error("couldn't delete user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrInternalServer)
	}

	h.logger.Info("user deleted", zap.Int32("user_id", int32(id)))
	return c.SendStatus(fiber.StatusNoContent)
}
