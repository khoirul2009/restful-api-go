package main

import (
	"learn-go-fiber/databases"
	"learn-go-fiber/internal/controllers"
	"learn-go-fiber/internal/repositories"
	"learn-go-fiber/internal/routes"
	"learn-go-fiber/internal/services"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {

	db := databases.InitDB()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	app := fiber.New()

	routes.SetupRoutes(app, userController)

	app.Listen(":3000")
}
