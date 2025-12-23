package service

import (
	"expenseTracker/internal/category/dto"
	"expenseTracker/internal/category/mapper"
	"expenseTracker/internal/category/repository"
	"expenseTracker/internal/user/service"
	"expenseTracker/pkg/errors"

	"github.com/google/uuid"
)

type CategoryService interface {
	Create(dto dto.CreateCategoryDTO) (dto.CategoryDTO, error)
	GetAll() ([]dto.CategoryDTO, error)
	GetByID(id uuid.UUID) (dto.CategoryDTO, error)
	Update(id uuid.UUID, dto dto.UpdateCategoryDTO) (dto.CategoryDTO, error)
	Delete(id uuid.UUID) error
	CategoryByUser(id uuid.UUID) ([]dto.CategoryDTO, error)
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
	userService        service.UserService
}

func NewCategoryService(categoryRepository repository.CategoryRepository, userService service.UserService) CategoryService {
	return &categoryService{categoryRepository, userService}
}

func (s *categoryService) Create(d dto.CreateCategoryDTO) (dto.CategoryDTO, error) {

	_, err := s.userService.GetByID(d.UserID.String())
	if err != nil {
		return dto.CategoryDTO{}, &errors.UserNotFoundError{d.UserID.String()}
	}

	category := mapper.DTOToCategory(&d)

	createdCategory, err := s.categoryRepository.Create(category)
	if err != nil {
		return dto.CategoryDTO{}, err
	}

	resultDTO := mapper.CategoryToDTO(createdCategory)
	return resultDTO, nil

}

func (s *categoryService) GetAll() ([]dto.CategoryDTO, error) {

	categoryList, err := s.categoryRepository.GetAll()
	if err != nil {
		return []dto.CategoryDTO{}, err
	}
	return mapper.CategoriesToDTO(categoryList), nil
}

func (s *categoryService) GetByID(id uuid.UUID) (dto.CategoryDTO, error) {

	category, err := s.categoryRepository.GetByID(id)
	if err != nil {
		return dto.CategoryDTO{}, err
	}
	return mapper.CategoryToDTO(category), nil
}

func (s *categoryService) Update(id uuid.UUID, update dto.UpdateCategoryDTO) (dto.CategoryDTO, error) {
	category, err := s.categoryRepository.GetByID(id)
	if err != nil {
		return dto.CategoryDTO{}, err
	}

	mapper.UpdateDtoToCategory(update, category)

	updatedCategory, err := s.categoryRepository.Update(category)
	if err != nil {
		return dto.CategoryDTO{}, err
	}

	return mapper.CategoryToDTO(updatedCategory), nil
}

func (s *categoryService) Delete(id uuid.UUID) error {

	err := s.categoryRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *categoryService) CategoryByUser(id uuid.UUID) ([]dto.CategoryDTO, error) {

	categories, err := s.categoryRepository.CategoriesByUser(id)
	if err != nil {
		return nil, err
	}
	return mapper.CategoriesToDTO(categories), nil
}
