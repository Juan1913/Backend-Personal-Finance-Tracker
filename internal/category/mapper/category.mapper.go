package mapper

import (
	"apiGo/internal/category/dto"
	"apiGo/internal/category/model"
)

func CategoryToDTO(category *model.Category) dto.CategoryDTO {

	return dto.CategoryDTO{
		ID:          category.ID.String(),
		Name:        category.Name,
		Description: category.Description,
		UserID:      category.UserID.String(),
		Type:        string(category.Type),
		CreatedAt:   category.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   category.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func CategoriesToDTO(categories []model.Category) []dto.CategoryDTO {
	dtos := make([]dto.CategoryDTO, len(categories))
	for i, u := range categories {
		dtos[i] = CategoryToDTO(&u)
	}
	return dtos
}

func DTOToCategory(d *dto.CreateCategoryDTO) *model.Category {
	return &model.Category{
		Name:        d.Name,
		Description: d.Description,
		UserID:      d.UserID,
		Type:        model.CategoryType(d.Type),
	}
}
