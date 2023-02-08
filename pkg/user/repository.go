package user

import (
	"Bugs-Bunny/pkg/models"
	"github.com/google/uuid"
)

type Repository interface {
	GetUserByEmail(email string) (*models.User, error)
	ReplaceApiKey(UserId *uuid.UUID) (*models.ApiKey, error)
}
