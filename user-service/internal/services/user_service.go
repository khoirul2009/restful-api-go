// internal/services/user_service.go
package services

import (
	"learn-go-fiber/internal/models"
	"learn-go-fiber/internal/repositories"
)

// UserService handles user-related business logic
type UserService struct {
	UserRepo *repositories.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// GetUsers retrieves all users
func (s *UserService) GetUsers() ([]models.User, error) {
	return s.UserRepo.FetchAllUsers()
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.UserRepo.FetchUserByID(id)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.InsertUser(user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id string, user *models.User) error {
	return s.UserRepo.UpdateUser(id, user)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id string) error {
	return s.UserRepo.DeleteUser(id)
}
