package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"util"

	"service"

	"github.com/julienschmidt/httprouter"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
	}
)

// NewUserController creating a new UserController instance
// func NewUserController(s *mgo.Session) *UserController {
// 	return &UserController{s}
// }

// NewUserController creating a new UserController instance
func NewUserController() *UserController {
	return &UserController{}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Get a UserService instance
	us := service.NewUserService(util.GetSession())
	u, err := us.GetUser(id)
	// Verify id is ObjectId, otherwise bail
	if err {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Get a UserService instance
	us := service.NewUserService(util.GetSession())
	user, err := us.CreateUser(u)
	if err {
		w.WriteHeader(400)
		return
	}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(user)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Get a UserService instance
	us := service.NewUserService(util.GetSession())
	err := us.RemoveUser(id)
	// Verify id is ObjectId, otherwise bail
	if err {
		w.WriteHeader(400)
		return
	}

	// Write status
	w.WriteHeader(200)
}
