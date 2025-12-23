package mapper

import (
	"expenseTracker/internal/user/dto"
	"expenseTracker/internal/user/model"
)

func UserToDTO(user *model.User) dto.UserDTO {
	return dto.UserDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Role:  string(user.Role),
	}
}

func UsersToDTO(users []model.User) []dto.UserDTO {
	dtos := make([]dto.UserDTO, len(users))
	for i, u := range users {
		dtos[i] = UserToDTO(&u)
	}
	return dtos
}

func DTOToUser(d *dto.CreateUserDTO) *model.User {
	return &model.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
		Role:     model.Role(d.Role),
	}
}

func UpdateUserFromDTO(user *model.User, d *dto.UpdateUserDTO) {
	if d.Name != "" {
		user.Name = d.Name
	}
	if d.Email != "" {
		user.Email = d.Email
	}
	if d.Password != "" {
		user.Password = d.Password
	}
	if d.Role != "" {
		user.Role = model.Role(d.Role)
	}
}
