package models

// User data
type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firsname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
