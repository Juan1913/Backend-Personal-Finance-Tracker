package model

import (
	"time"

	"github.com/google/uuid"
)

type RecurringFrequency string

const (
	Monthly  RecurringFrequency = "monthly"
	Weekly   RecurringFrequency = "weekly"
	Biweekly RecurringFrequency = "biweekly"
	Yearly   RecurringFrequency = "yearly"
	Custom   RecurringFrequency = "custom"
)

type RecurringRule struct {
	ID          uuid.UUID          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID      uuid.UUID          `gorm:"type:uuid;not null" json:"user_id"`
	Name        string             `gorm:"size:100;not null" json:"name"`
	Type        string             `gorm:"type:varchar(10);not null" json:"type"` // income, expense
	AccountID   uuid.UUID          `gorm:"type:uuid;not null" json:"account_id"`
	CategoryID  uuid.UUID          `gorm:"type:uuid;not null" json:"category_id"`
	Amount      float64            `gorm:"type:decimal(14,2);not null" json:"amount"`
	Frequency   RecurringFrequency `gorm:"type:varchar(20);not null" json:"frequency"`
	StartDate   time.Time          `json:"start_date"`
	EndDate     *time.Time         `json:"end_date,omitempty"`
	NextRunDate time.Time          `json:"next_run_date"`
	Description string             `gorm:"type:text" json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
