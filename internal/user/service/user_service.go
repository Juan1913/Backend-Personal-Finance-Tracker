package service

import (
	"expenseTracker/internal/user/dto"
	userserrors "expenseTracker/internal/user/errors"
	"expenseTracker/internal/user/mapper"
	"expenseTracker/internal/user/repository"

	"github.com/google/uuid"
)

type UserService interface {
	Create(dto dto.CreateUserDTO) (dto.UserDTO, error)
	GetAll() ([]dto.UserDTO, error)
	GetByID(id string) (dto.UserDTO, error)
	Update(id string, dto dto.UpdateUserDTO) (dto.UserDTO, error)
	Delete(id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(d dto.CreateUserDTO) (dto.UserDTO, error) {
	existing, err := s.repo.FindByEmail(d.Email)
	if err == nil && existing != nil {
		return dto.UserDTO{}, &userserrors.EmailAlreadyExistsError{Email: d.Email}
	}
	user := mapper.DTOToUser(&d)
	err = s.repo.Create(user)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return mapper.UserToDTO(user), nil
}

func (s *userService) GetAll() ([]dto.UserDTO, error) {
	userList, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.UsersToDTO(userList), nil
}

func (s *userService) GetByID(id string) (dto.UserDTO, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return dto.UserDTO{}, err
	}
	user, err := s.repo.FindByID(uuidID)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return mapper.UserToDTO(user), nil
}

func (s *userService) Update(id string, d dto.UpdateUserDTO) (dto.UserDTO, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return dto.UserDTO{}, err
	}
	user, err := s.repo.FindByID(uuidID)
	if err != nil {
		return dto.UserDTO{}, err
	}
	//mapper para actualizar los campos del usuario
	mapper.UpdateUserFromDTO(user, &d)
	err = s.repo.Update(user)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return mapper.UserToDTO(user), nil
}

func (s *userService) Delete(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(uuidID)
}
