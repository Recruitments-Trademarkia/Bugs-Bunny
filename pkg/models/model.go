package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *sql.NullTime `gorm:"index"`
}
