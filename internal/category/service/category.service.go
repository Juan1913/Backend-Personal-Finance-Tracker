package service

import (
	"apiGo/internal/category/dto"
	"apiGo/internal/category/repository"
)

type CategoryService interface {
	Create(dto dto.CategoryDTO) (dto.CategoryDTO, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}
