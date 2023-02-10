package user

import (
	"Bugs-Bunny/pkg/models"

	"github.com/google/uuid"
)

type Service interface {
	GetUserByEmail(email string) (*models.User, error)
	ReplaceApiKey(UserId *uuid.UUID) (*models.ApiKey, error)
}

type userSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &userSvc{repo: r}
}

func (u *userSvc) GetUserByEmail(email string) (*models.User, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userSvc) ReplaceApiKey(UserId *uuid.UUID) (*models.ApiKey, error) {
	apiKey, err := u.repo.ReplaceApiKey(UserId)
	if err != nil {
		return nil, err
	}
	return apiKey, nil
}
