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
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil

	//return nil, nil
}

// ReplaceApiKey replaces the api key for a user
func (r *repo) ReplaceApiKey(UserId *uuid.UUID) (*models.ApiKey, error) {
	// TODO: Implement ReplaceApiKey
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

	//return nil, nil
}
