package services

import (
	"github.com/invent360/07-go-app/utils"

	"github.com/invent360/07-go-app/models"
	"github.com/invent360/07-go-app/repositories"
)

// GetUser Fetches user by Id
func GetUser(userID int64) (*models.User, *utils.ApplicationError) {
	return repositories.GetUser(userID)
}
