package model

import (
	"github.com/google/uuid"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type AuthUser struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email    string    `gorm:"size:100;unique;not null" json:"email"`
	Password string    `gorm:"size:100;not null" json:"password"`
	Role     Role      `gorm:"size:20;not null" json:"role"`
}
