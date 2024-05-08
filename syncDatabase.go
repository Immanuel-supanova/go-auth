package goauth

import (
	"github.com/immanuel-supanova/go-auth/database"
	"github.com/immanuel-supanova/go-auth/models"
)

func SyncDatabase() {
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Log{})

}
