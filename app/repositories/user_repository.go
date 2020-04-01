package repositories

import (
	"fmt"
	"net/http"
	"utils"

	"github.com/07-go-app/models"
)

var (
	users = map[int64]*models.User{
		123: {ID: 1, FirstName: "Fed", LastName: "Sack", Email: "fed.sack@gmail.com"},
	}
)

//GetUser returns a new models.user.
func GetUser(userID int64) (*models.User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User with ID %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "Not found",
	}
}
