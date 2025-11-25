package model

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionTag struct {
	TransactionID uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"transaction_id"`
	TagID         uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"tag_id"`
}
