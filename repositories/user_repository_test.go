package repositories

import "testing"

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)
	if user != nil {
		t.Error("we were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("we were expecting an error when user id is 0")
	}
}
