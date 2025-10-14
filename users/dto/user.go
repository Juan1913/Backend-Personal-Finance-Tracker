package dto

type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UpdateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
