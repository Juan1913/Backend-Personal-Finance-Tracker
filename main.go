package main

import (
	"apiGo/internal/account"
	accountHandler "apiGo/internal/account/handler"
	accountRepository "apiGo/internal/account/repository"
	accountService "apiGo/internal/account/service"
	authHandler "apiGo/internal/auth/handler"
	authRepository "apiGo/internal/auth/repository"
	authRoutes "apiGo/internal/auth/routes"
	authService "apiGo/internal/auth/service"
	"apiGo/internal/bootstrap"
	"apiGo/internal/user"
	usersHandler "apiGo/internal/user/handler"
	usersRepository "apiGo/internal/user/repository"
	usersService "apiGo/internal/user/service"
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

	// Inyección de dependencias para user
	userRepo := usersRepository.NewUserRepository(db)
	userService := usersService.NewUserService(userRepo)
	userHandler := usersHandler.NewUserHandler(userService)

	// Inyección de dependencias para accounts
	accountRepo := accountRepository.NewAccountRepository(db)
	accountServ := accountService.NewAccountService(accountRepo)
	accountHand := accountHandler.NewAccountHandler(accountServ)

	// Auth module
	authRepo := authRepository.NewAuthRepository(db)
	authServ := authService.NewAuthService(authRepo)
	authHand := authHandler.NewAuthHandler(authServ)

	r := gin.Default()
	errors.RegisterUserRoutes(r, userHandler)
	account.RegisterAccountRoutes(r, accountHand)
	authRoutes.RegisterAuthRoutes(r, authHand)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
