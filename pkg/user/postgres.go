package user

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

// GetUserByEmail returns a user by email
func (r *repo) GetUserByEmail(email string) (*models.User, error) {
	// TODO: Implement GetUserByEmail
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil

	//return nil, nil
}

// ReplaceApiKey replaces the api key for a user
func (r *repo) ReplaceApiKey(UserId *uuid.UUID) (*models.ApiKey, error) {
	/* TODO: Implement ReplaceApiKey
	apiKeyUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// Create a new API key object
	apiKey := &models.ApiKey{
		UserId:    &apiKeyUUID,
		BaseModel: models.BaseModel{},
	}

	// Save the API key to the database
	if err := r.DB.Save(apiKey).Error; err != nil {
		return nil, err
	}

	return apiKey, nil


	//return nil, nil*/
	apiKeyUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// Fetch the existing API key
	var apiKey models.ApiKey
	err = r.DB.Where("user_id = ?", UserId).First(&apiKey).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new API key object
			apiKey = models.ApiKey{
				BaseModel: apiKey.BaseModel,
				UserId:    &apiKeyUUID,
			}

			// Save the new API key to the database
			if err := r.DB.Create(&apiKey).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// Update the existing API key
		apiKey.UserId = &apiKeyUUID

		// Save the updated API key to the database
		if err := r.DB.Save(&apiKey).Error; err != nil {
			return nil, err
		}
	}

	return &apiKey, nil

}
