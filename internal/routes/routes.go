package routes

import (
	"learn-go-fiber/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes defines the API routes and injects the controller dependencies
func SetupRoutes(app *fiber.App, userController *controllers.UserController) {
	// Basic route to check the server
	app.Get("/", userController.Home)

	// User-related routes
	userRoutes := app.Group("/users")
	userRoutes.Get("/", userController.GetAllUsers)      // GET /users - Fetch all users
	userRoutes.Get("/:id", userController.GetUser)       // GET /users/:id - Fetch user by ID
	userRoutes.Post("/", userController.CreateUser)      // POST /users - Create a new user
	userRoutes.Put("/:id", userController.UpdateUser)    // PUT /users/:id - Update user by ID
	userRoutes.Delete("/:id", userController.DeleteUser) // DELETE /users/:id - Delete user by ID
}
