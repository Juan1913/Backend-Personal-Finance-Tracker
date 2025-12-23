package repository

import (
	"expenseTracker/internal/category/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	GetAll() ([]model.Category, error)
	GetByID(id uuid.UUID) (*model.Category, error)
	Update(category *model.Category) (*model.Category, error)
	Delete(id uuid.UUID) error
	CategoriesByUser(id uuid.UUID) ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}


func (r *categoryRepository) Create(category *model.Category) (*model.Category, error) {

	err := r.db.Create(category).Error
	return category, err
}

func (r *categoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByID(id uuid.UUID) (*model.Category, error) {

	var category model.Category
	err := r.db.Find(&category, "id = ?", id).Error
	return &category, err

}

func (r *categoryRepository) Update(category *model.Category) (*model.Category, error) {
	err := r.db.Save(category).Error
	return category, err
}

func (r *categoryRepository) Delete(id uuid.UUID) error {

	var category model.Category
	err := r.db.Delete(&category, "id = ?", id).Error

	return err
}

func (r *categoryRepository) CategoriesByUser(id uuid.UUID) ([]model.Category, error) {

	var categoriesByUser []model.Category
	err := r.db.Where("user_id = ?", id).Find(&categoriesByUser).Error

	return categoriesByUser, err
}
