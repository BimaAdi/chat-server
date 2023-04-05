package tasks

import (
	"time"

	"github.com/BimaAdi/chat-server/models"
	"github.com/BimaAdi/chat-server/repository"
	"github.com/BimaAdi/chat-server/settings"
)

func CreateSuperUser(envPath string, email string, username string, password string) {
	// Initialize environtment variable
	settings.InitiateSettings(envPath)

	// Initiate Database connection
	models.Initiate()

	now := time.Now()
	repository.CreateUser(models.DBConn, username, email, password, true, true, now, &now)
}
