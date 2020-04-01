package services

import (
	"github.com/invent360/07-go-app/app/utils"

	"github.com/invent360/07-go-app/app/models"
	"github.com/invent360/07-go-app/app/repositories"
)

// GetUser Fetches user by Id
func GetUser(userID int64) (*models.User, *utils.ApplicationError) {
	return repositories.GetUser(userID)
}
