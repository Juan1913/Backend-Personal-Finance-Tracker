package repository

import (
	"apiGo/internal/category/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	GetAll([]model.Category) ([]model.Category, error)
	GetByID(id uuid.UUID) (*model.Category, error)
	Delete(id uuid.UUID) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {

}
