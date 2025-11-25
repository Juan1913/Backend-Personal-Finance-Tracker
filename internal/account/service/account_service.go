package service

import (
	"apiGo/internal/account/dto"
	"apiGo/internal/account/mapper"
	"apiGo/internal/account/repository"
	userModel "apiGo/internal/user/model"

	"github.com/google/uuid"
)

type AccountService interface {
	Create(dto dto.CreateAccountDTO) (dto.AccountResponseDTO, error)
	GetAll() ([]dto.AccountResponseDTO, error)
	GetByID(id string) (dto.AccountResponseDTO, error)
	GetByUserID(userID string) ([]dto.AccountResponseDTO, error)
	Update(id string, dto dto.UpdateAccountDTO) (dto.AccountResponseDTO, error)
	Delete(id string) error
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo}
}

func (s *accountService) Create(d dto.CreateAccountDTO) (dto.AccountResponseDTO, error) {
	account, err := mapper.DTOToAccount(&d)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	err = s.repo.Create(account)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	// Obtener la cuenta completa con el usuario
	createdAccount, user, err := s.repo.FindByID(account.ID)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	return mapper.AccountToResponseDTO(createdAccount, user), nil
}

func (s *accountService) GetAll() ([]dto.AccountResponseDTO, error) {
	accounts, users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return mapper.AccountsToResponseDTO(accounts, users), nil
}

func (s *accountService) GetByID(id string) (dto.AccountResponseDTO, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	account, user, err := s.repo.FindByID(uuidID)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	return mapper.AccountToResponseDTO(account, user), nil
}

func (s *accountService) GetByUserID(userID string) ([]dto.AccountResponseDTO, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	accounts, user, err := s.repo.FindByUserID(uuidUserID)
	if err != nil {
		return nil, err
	}

	// Crear un slice con el usuario repetido para cada cuenta
	users := make([]userModel.User, len(accounts))
	for i := range accounts {
		users[i] = *user
	}

	return mapper.AccountsToResponseDTO(accounts, users), nil
}

func (s *accountService) Update(id string, d dto.UpdateAccountDTO) (dto.AccountResponseDTO, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	account, user, err := s.repo.FindByID(uuidID)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	mapper.UpdateAccountFromDTO(account, &d)

	err = s.repo.Update(account)
	if err != nil {
		return dto.AccountResponseDTO{}, err
	}

	return mapper.AccountToResponseDTO(account, user), nil
}

func (s *accountService) Delete(id string) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(uuidID)
}
