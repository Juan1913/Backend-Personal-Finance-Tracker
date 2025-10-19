package bootstrap

import (
	accountModel "apiGo/account/model"
	usersModel "apiGo/users/model"
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
