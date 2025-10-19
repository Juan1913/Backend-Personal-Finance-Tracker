package dto

import "apiGo/users/dto"

type AccountDTO struct {
	ID                 string `json:"id"`
	AccountName        string `json:"account_name"`
	AccountDescription string `json:"account_description,omitempty"`
	UserID             string `json:"user_id"`
}

type AccountResponseDTO struct {
	ID                 string      `json:"id"`
	AccountName        string      `json:"account_name"`
	AccountDescription string      `json:"account_description,omitempty"`
	UserID             string      `json:"user_id"`
	User               dto.UserDTO `json:"user"`
}

type CreateAccountDTO struct {
	AccountName        string `json:"account_name" binding:"required"`
	AccountDescription string `json:"account_description"`
	UserID             string `json:"user_id" binding:"required"`
}

type UpdateAccountDTO struct {
	AccountName        string `json:"account_name"`
	AccountDescription string `json:"account_description"`
}
