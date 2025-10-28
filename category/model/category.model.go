package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryType string

const (
	IncomeCategory  CategoryType = "income"
	ExpenseCategory CategoryType = "expense"
)

type Category struct {
	ID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string       `gorm:"size:100;not null" json:"name"`
	Description string       `gorm:"size:255" json:"description"`
	UserID      uuid.UUID    `gorm:"type:uuid;not null" json:"user_id"`
	Type        CategoryType `gorm:"type:varchar(10);not null" json:"type"`
	ParentID    *uuid.UUID   `gorm:"type:uuid" json:"parent_id,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
