package api_keys

import "github.com/google/uuid"

type Repository interface {
	GetApiKeyByUserId(userId *uuid.UUID) (*uuid.UUID, error)
	ApiAuthCheck(ApiKey *uuid.UUID) (bool, error)
}
