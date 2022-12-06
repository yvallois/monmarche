package models

import (
	"time"

	"github.com/google/uuid"
)

type ModelBase struct {
	UUID      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null" json:"updated_at"`
}
