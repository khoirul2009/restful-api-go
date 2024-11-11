// internal/controllers/user_controller.go
package controllers

import (
	"learn-go-fiber/internal/models"
	"learn-go-fiber/internal/request"
	"learn-go-fiber/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// UserController handles user-related HTTP requests
type UserController struct {
	UserService *services.UserService
}

// NewUserController creates a new instance of UserController
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

// Home handles the home route
func (uc *UserController) Home(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Go Fiber API!")
}

// GetAllUsers handles GET requests to retrieve all users
func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch users"})
	}
	return c.JSON(users)
}

// GetUser handles GET requests to retrieve a user by ID
func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := uc.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// CreateUser handles POST requests to create a new user
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	// Bind request ke CreateUserRequest

	var req = &request.CreateUserRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	// Validasi input menggunakan go-playground/validator
	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate UUID v4 untuk id
	userID := uuid.New().String()

	// Map dari request ke model User
	user := &models.User{
		ID:       userID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // Password hashing sebaiknya di handle di service layer
	}

	// Panggil service untuk create user
	err := uc.UserService.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Response sukses
	return c.Status(fiber.StatusCreated).JSON(user)
}

// UpdateUser handles PUT requests to update an existing user
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	err := uc.UserService.UpdateUser(id, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(user)
}

// DeleteUser handles DELETE requests to delete a user by ID
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := uc.UserService.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
