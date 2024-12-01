package services

import (
	"fmt"
	"learn-go-fiber/internal/exceptions"
	"learn-go-fiber/internal/models"
	"learn-go-fiber/internal/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	SessionRepo *repositories.SessionRepository
	UserRepo    *repositories.UserRepository
}

func NewAuthService(sessionRepo *repositories.SessionRepository, userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{SessionRepo: sessionRepo, UserRepo: userRepo}
}

func (as *AuthService) CreateSession(User *models.User, DeviceInfo string) (models.SessionToken, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	var sessionToken models.SessionToken

	// Set token claims
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = map[string]interface{}{
		"user_id": User.ID,
		"email":   User.Email,
	}
	claims["exp"] = time.Now().Add(time.Hour * 168).Unix()

	secretKey := os.Getenv("JWT_SECRET")

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return models.SessionToken{}, exceptions.NewHttpException(500, "Failed to generate token")
	}

	sessionToken.ID = uuid.New().String()
	sessionToken.Token = tokenString
	sessionToken.UserID = User.ID
	sessionToken.ExpiredAt = time.Now()
	sessionToken.CreatedAt = time.Now()
	sessionToken.DeviceInfo = DeviceInfo
	sessionToken.User = *User

	err = as.SessionRepo.StoreToken(&sessionToken)

	if err != nil {
		return models.SessionToken{}, err
	}

	return sessionToken, nil
}

func (as *AuthService) CurrentSession(token string) (*models.SessionToken, *models.User, error) {
	secretKey := os.Getenv("JWT_SECRET")

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, nil, exceptions.NewHttpException(401, "Invalied token")
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userClaims := claims["user"].(map[string]interface{})
		userID := userClaims["user_id"].(string)
		exp := int64(claims["exp"].(float64)) // Claims store numbers as float64

		// Retrieve user data from the database

		user, err := as.UserRepo.FetchUserByID(userID)
		if err != nil {
			return nil, nil, exceptions.NewHttpException(404, "user not found")
		}

		// Return session details and user data
		sessionToken := &models.SessionToken{
			UserID:     userID,
			Token:      token,
			ExpiredAt:  time.Unix(exp, 0),
			DeviceInfo: claims["device_info"].(string), // Optional, if included in JWT
		}

		return sessionToken, user, nil
	}

	return nil, nil, exceptions.NewHttpException(401, "Invalid token")
}
