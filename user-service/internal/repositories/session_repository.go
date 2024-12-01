package repositories

import (
	"errors"
	"learn-go-fiber/internal/exceptions"
	"learn-go-fiber/internal/models"

	"gorm.io/gorm"
)

type SessionRepository struct {
	DB *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{DB: db}
}

func (sr *SessionRepository) GetToken(token string) (*models.SessionToken, error) {
	var session models.SessionToken
	if err := sr.DB.Select("token ").First(&session, "token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.NewHttpException(500, "failed query data :"+err.Error())
		}
		return nil, exceptions.NewHttpException(500, "failed query data :"+err.Error())
	}
	return &session, nil
}

func (sr *SessionRepository) StoreToken(sessionToken *models.SessionToken) error {

	if err := sr.DB.Create(&sessionToken).Error; err != nil {
		return exceptions.NewHttpException(500, "failed to store token")
	}
	return nil
}

func (sr *SessionRepository) DeleteToken(token string) error {
	var sessionToken models.SessionToken

	if err := sr.DB.Where("token = ?", token).First(&sessionToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return exceptions.NewHttpException(404, "token not found")
		}
		return exceptions.NewHttpException(500, "Failed to query database: "+err.Error())
	}

	// Hapus token
	if err := sr.DB.Delete(&sessionToken).Error; err != nil {
		return exceptions.NewHttpException(500, "failed to delete token: "+err.Error())
	}

	return nil
}
