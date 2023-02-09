package api_keys

import (
	"Bugs-Bunny/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

// GetApiKeyByUserId returns the api key for a given user id
func (r *repo) GetApiKeyByUserId(userId *uuid.UUID) (*uuid.UUID, error) {
	// TODO: implement this
	var apiKey uuid.UUID
	err := r.DB.Where("user_id = ?", userId).First(&apiKey).Error
	if err != nil {
		return nil, err
	}

	return &apiKey, nil

	//return nil, nil
}

// ApiAuthCheck checks if the api key is valid for the given user id
func (r *repo) ApiAuthCheck(ApiKeys *uuid.UUID) (bool, error) {
	// TODO: implement this
	var count int64
	err := r.DB.Model(&models.ApiKey{}).Where("api_key = ?", ApiKeys).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
