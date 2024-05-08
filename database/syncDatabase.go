package database

import "github.com/immanuel-supanova/go-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Log{})

}
