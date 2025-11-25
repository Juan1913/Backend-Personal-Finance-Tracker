package dto

import "github.com/google/uuid"

type CategoryDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Type        string `json:"type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateCategoryDTO struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=income expense"`
}

type UpdateCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type" binding:"omitempty,oneof=income expense"`
}

type CategoryResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	Type        string `json:"type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
