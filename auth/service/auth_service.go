package service

import (
	"apiGo/auth/dto"
	"apiGo/auth/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecret")

type AuthService interface {
	Login(dto dto.LoginDTO) (dto.AuthResponseDTO, error)
	Register(dto dto.RegisterDTO) (dto.AuthResponseDTO, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(d dto.LoginDTO) (dto.AuthResponseDTO, error) {
	user, err := s.repo.FindByEmail(d.Email)
	if err != nil || user == nil || user.Password != d.Password {
		return dto.AuthResponseDTO{}, err
	}
	token := generateToken(user.Email, string(user.Role))
	return dto.AuthResponseDTO{Token: token, Role: string(user.Role), User: user}, nil
}

func (s *authService) Register(d dto.RegisterDTO) (dto.AuthResponseDTO, error) {
	// El registro de usuarios se maneja en el módulo users, aquí solo autenticación
	return dto.AuthResponseDTO{}, nil
}

func generateToken(email, role string) string {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(jwtSecret)
	return t
}
