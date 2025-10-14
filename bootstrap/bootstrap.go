package bootstrap

import (
	usersModel "apiGo/users/model"
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(

		&usersModel.User{}); err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}
}
