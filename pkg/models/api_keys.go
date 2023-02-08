package models

import "github.com/google/uuid"

type ApiKey struct {
	BaseModel
	UserId *uuid.UUID `gorm:"type:uuid;not null"`
}
