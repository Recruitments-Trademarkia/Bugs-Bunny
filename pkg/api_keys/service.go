package api_keys

import "github.com/google/uuid"

type Service interface {
	GetApiKeyByUserId(userId *uuid.UUID) (*uuid.UUID, error)
	ApiAuthCheck(ApiKey *uuid.UUID) (bool, error)
}

type apiKeysSvc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &apiKeysSvc{repo: r}
}

func (s *apiKeysSvc) GetApiKeyByUserId(userId *uuid.UUID) (*uuid.UUID, error) {
	return s.repo.GetApiKeyByUserId(userId)
}

func (s *apiKeysSvc) ApiAuthCheck(ApiKey *uuid.UUID) (bool, error) {
	return s.repo.ApiAuthCheck(ApiKey)
}
