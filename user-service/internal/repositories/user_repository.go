// internal/repositories/user_repository.go
package repositories

import (
	"errors"
	"learn-go-fiber/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository represents the database connection
// UserRepository represents the database connection
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// FetchAllUsers fetches all users from the database
func (r *UserRepository) FetchAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Select("id", "name", "email").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FetchUserByID fetches a user by ID from the database
func (r *UserRepository) FetchUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.DB.Select("id", "name", "email").First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// InsertUser inserts a new user into the database
func (r *UserRepository) InsertUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(id string, updatedUser *models.User) error {
	var user models.User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	// Update the user fields
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email

	if err := r.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database
func (r *UserRepository) DeleteUser(id string) error {
	if err := r.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
