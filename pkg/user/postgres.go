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

	return nil, nil
}

// ReplaceApiKey replaces the api key for a user
func (r *repo) ReplaceApiKey(UserId *uuid.UUID) (*models.ApiKey, error) {
	// TODO: Implement ReplaceApiKey

	return nil, nil
}
