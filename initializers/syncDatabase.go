package initializers

import (
	"log"

	"github.com/NerdyNarayan/go-jwt/models"
)

func SyncDatabase() {
	if DB == nil {
		log.Fatal("Database not connected")
	}
	DB.AutoMigrate(&models.User{})

}
