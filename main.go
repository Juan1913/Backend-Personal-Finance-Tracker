package main

import (
	"expenseTracker/internal/bootstrap"
	"expenseTracker/internal/config"
	"expenseTracker/internal/factory"
	"expenseTracker/internal/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dsn := cfg.DB.DSN()
	// Configuración de la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// Migración automática
	bootstrap.AutoMigrate(db)

	// Inyección de dependencias con AppFactory
	deps := factory.NewAppDependencies(db)

	r := gin.Default()
	routes.RegisterUserRoutes(r, deps.UserHandler)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
