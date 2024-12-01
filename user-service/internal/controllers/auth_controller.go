package controllers

import (
	"learn-go-fiber/internal/exceptions"
	"learn-go-fiber/internal/request"
	"learn-go-fiber/internal/services"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

func NewAuthController(authService *services.AuthService, userService *services.UserService) *AuthController {
	return &AuthController{AuthService: authService, UserService: userService}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var req request.LoginRequest
	var validate = validator.New()

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := ac.UserService.UserRepo.FetchUserByEmail(req.Email)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalied credentials"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalied credentials"})
	}

	session, err := ac.AuthService.CreateSession(user, string(c.Context().UserAgent()))

	if err != nil {
		if httpErr, ok := err.(*exceptions.HttpException); ok {
			return c.Status(httpErr.Code).JSON(fiber.Map{"error": httpErr.Message})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	secretKey := os.Getenv("JWT_SECRET")

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return exceptions.NewHttpException(500, "Failed to generate token")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": tokenString,
		"data":         session,
	})

}
