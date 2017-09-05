package main

import (
	// Standard library packages
	"net/http"

	// Third party packages
	"controllers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	//uc := controllers.NewUserController()

	// Get a UserController instance
	uc := controllers.NewUserController()

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	r.POST("/user", uc.CreateUser)

	r.DELETE("/user/:id", uc.RemoveUser)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}
