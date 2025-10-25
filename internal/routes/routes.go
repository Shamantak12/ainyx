package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"user-api/internal/handler"
	"user-api/internal/middleware"
	"user-api/internal/service"
)

func SetupRoutes(app *fiber.App, userService service.UserService, logger *zap.Logger) {
	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.LoggerMiddleware(logger))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "User API is running",
		})
	})

	api := app.Group("/api/v1")

	userHandler := handler.NewUserHandler(userService, logger)

	users := api.Group("/users")
	users.Post("/", userHandler.CreateUser)
	users.Get("/", userHandler.ListUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}
