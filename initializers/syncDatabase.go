package initializers

import "github.com/GingFreecss2/Go-jwt-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
