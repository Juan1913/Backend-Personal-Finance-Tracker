package mapper

import (
	"apiGo/internal/account/dto"
	"apiGo/internal/account/model"
	"apiGo/internal/user/mapper"
	userModel "apiGo/internal/user/model"

	"github.com/google/uuid"
)

// AccountToResponseDTO convierte Account y User a AccountResponseDTO (para respuestas)
func AccountToResponseDTO(account *model.Account, user *userModel.User) dto.AccountResponseDTO {
	return dto.AccountResponseDTO{
		ID:                 account.ID.String(),
		AccountName:        account.AccountName,
		AccountDescription: account.AccountDescription,
		UserID:             account.UserID.String(),
		User:               mapper.UserToDTO(user),
	}
}

// AccountsToResponseDTO convierte slice de Account y Users a slice de AccountResponseDTO
func AccountsToResponseDTO(accounts []model.Account, users []userModel.User) []dto.AccountResponseDTO {
	dtos := make([]dto.AccountResponseDTO, len(accounts))
	userMap := make(map[string]*userModel.User)

	// Crear un mapa de usuarios por ID para búsqueda rápida
	for i := range users {
		userMap[users[i].ID.String()] = &users[i]
	}

	for i, a := range accounts {
		user := userMap[a.UserID.String()]
		dtos[i] = AccountToResponseDTO(&a, user)
	}
	return dtos
}

func DTOToAccount(d *dto.CreateAccountDTO) (*model.Account, error) {
	userID, err := uuid.Parse(d.UserID)
	if err != nil {
		return nil, err
	}

	return &model.Account{
		AccountName:        d.AccountName,
		AccountDescription: d.AccountDescription,
		UserID:             userID,
	}, nil
}

func UpdateAccountFromDTO(account *model.Account, d *dto.UpdateAccountDTO) {
	if d.AccountName != "" {
		account.AccountName = d.AccountName
	}
	if d.AccountDescription != "" {
		account.AccountDescription = d.AccountDescription
	}
}
