package main

import (
	authHandler "apiGo/auth/handler"
	authRepository "apiGo/auth/repository"
	authRoutes "apiGo/auth/routes"
	authService "apiGo/auth/service"
	"apiGo/bootstrap"
	"apiGo/users"
	usersHandler "apiGo/users/handler"
	usersRepository "apiGo/users/repository"
	usersService "apiGo/users/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configuración de la base de datos
	dsn := "host=localhost user=user password=password dbname=apiGo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// Migración automática
	bootstrap.AutoMigrate(db)

	// Inyección de dependencias para users
	repo := usersRepository.NewUserRepository(db)
	service := usersService.NewUserService(repo)
	handler := usersHandler.NewUserHandler(service)

	// Auth module
	authRepo := authRepository.NewAuthRepository(db)
	authServ := authService.NewAuthService(authRepo)
	authHand := authHandler.NewAuthHandler(authServ)

	r := gin.Default()
	users.RegisterUserRoutes(r, handler)
	authRoutes.RegisterAuthRoutes(r, authHand)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
