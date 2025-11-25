package model

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

type TransactionStatus string

const (
	IncomeTransaction   TransactionType = "income"
	ExpenseTransaction  TransactionType = "expense"
	TransferTransaction TransactionType = "transfer"

	PendingStatus    TransactionStatus = "pending"
	ClearedStatus    TransactionStatus = "cleared"
	ReconciledStatus TransactionStatus = "reconciled"
)

type Transaction struct {
	ID               uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID           uuid.UUID         `gorm:"type:uuid;not null" json:"user_id"`
	Date             time.Time         `json:"date"`
	Type             TransactionType   `gorm:"type:varchar(10);not null" json:"type"`
	AccountID        uuid.UUID         `gorm:"type:uuid;not null" json:"account_id"`
	CounterAccountID *uuid.UUID        `gorm:"type:uuid" json:"counter_account_id,omitempty"`
	CategoryID       *uuid.UUID        `gorm:"type:uuid" json:"category_id,omitempty"`
	MerchantName     *string           `gorm:"size:100" json:"merchant_name,omitempty"`
	Amount           float64           `gorm:"type:decimal(14,2);not null" json:"amount"`
	Currency         string            `gorm:"size:10;not null" json:"currency"`
	Description      string            `gorm:"type:text" json:"description"`
	Notes            string            `gorm:"type:text" json:"notes"`
	Status           TransactionStatus `gorm:"type:varchar(12);not null" json:"status"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}
