package bootstrap

import (
	accountModel "expenseTracker/internal/account/model"
	usersModel "expenseTracker/internal/user/model"
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&usersModel.User{},
		&accountModel.Account{}); err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}
}
