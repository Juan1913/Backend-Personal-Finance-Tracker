package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name     string    `gorm:"size:100;not null" json:"name"`
	Email    string    `gorm:"size:100;unique;not null" json:"email"`
	Password string    `gorm:"size:100;not null" json:"password"`
	Role     Role      `gorm:"size:20;not null" json:"role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
