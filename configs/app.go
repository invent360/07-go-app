package configs

import (
	"net/http"

	"github.com/invent360/07-go-app/controllers"
)

// StartApp starts new go app
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
