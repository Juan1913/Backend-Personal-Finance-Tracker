package factory

import (
	"expenseTracker/internal/user/handler"
	"expenseTracker/internal/user/repository"
	"expenseTracker/internal/user/service"

	"gorm.io/gorm"
)

type AppDependencies struct {
	UserHandler *handler.UserHandler
}

func NewAppDependencies(db *gorm.DB) *AppDependencies {

	//user
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//categories

	return &AppDependencies{
		UserHandler: userHandler,
	}
}
