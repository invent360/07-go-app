package services

import (
	"utils"

	"github.com/07-go-app/models"
	"github.com/07-go-app/repositories"
)

// GetUser Fetches user by Id
func GetUser(userID int64) (*models.User, *utils.ApplicationError) {
	return repositories.GetUser(userID)
}