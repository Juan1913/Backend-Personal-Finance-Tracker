package model

import (
	"apiGo/users/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	AccountName        string    `gorm:"size:100;not null" json:"account_name"`
	AccountDescription string    `gorm:"size:255" json:"account_description"`
	UserID             uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
}

type AccountWithUser struct {
	ID                 uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	AccountName        string     `gorm:"size:100;not null" json:"account_name"`
	AccountDescription string     `gorm:"size:255" json:"account_description"`
	UserID             uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	User               model.User `gorm:"foreignKey:UserID" json:"user"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
