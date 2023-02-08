package api_keys

import (
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
	return nil, nil
}

// ApiAuthCheck checks if the api key is valid for the given user id
func (r *repo) ApiAuthCheck(ApiKey *uuid.UUID) (bool, error) {
	// TODO: implement this
	return false, nil
}
