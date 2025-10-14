package service

import (
	"apiGo/users/dto"
	"apiGo/users/model"
	"apiGo/users/repository"

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
	user := model.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
		Role:     model.Role(d.Role),
	}
	err := s.repo.Create(&user)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return dto.UserDTO{ID: user.ID.String(), Name: user.Name, Email: user.Email, Role: string(user.Role)}, nil
}

func (s *userService) GetAll() ([]dto.UserDTO, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	dtos := make([]dto.UserDTO, len(users))
	for i, u := range users {
		dtos[i] = dto.UserDTO{ID: u.ID.String(), Name: u.Name, Email: u.Email, Role: string(u.Role)}
	}
	return dtos, nil
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
	return dto.UserDTO{ID: user.ID.String(), Name: user.Name, Email: user.Email, Role: string(user.Role)}, nil
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
	if d.Name != "" {
		user.Name = d.Name
	}
	if d.Email != "" {
		user.Email = d.Email
	}
	err = s.repo.Update(user)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return dto.UserDTO{ID: user.ID.String(), Name: user.Name, Email: user.Email}, nil
}

func (s *userService) Delete(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(uuidID)
}
