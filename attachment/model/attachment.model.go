package model

import (
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	TransactionID uuid.UUID `gorm:"type:uuid;not null" json:"transaction_id"`
	FileName      string    `gorm:"size:255;not null" json:"file_name"`
	FileURL       string    `gorm:"size:255;not null" json:"file_url"`
	MimeType      string    `gorm:"size:100;not null" json:"mime_type"`
	CreatedAt     time.Time `json:"created_at"`
}
