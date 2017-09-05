// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"models"

// 	"github.com/julienschmidt/httprouter"
// )

// func main() {

// 	r := httprouter.New()

// 	r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		u := models.User{
// 			Name:   "Puneet Joshi",
// 			Gender: "male",
// 			Age:    22,
// 			Id:     p.ByName("Id"),
// 		}
// 		// Marshal provided interface into JSON structure
// 		uj, _ := json.Marshal(u)

// 		// Write content-type, statuscode, payload
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(200)
// 		fmt.Fprintf(w, "%s", uj)
// 	})

// 	r.POST("/user", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 		// Stub an user to be populated from the body
// 		u := models.User{}

// 		// Populate the user data
// 		json.NewDecoder(r.Body).Decode(&u)

// 		// Add an Id
// 		u.Id = "foo"

// 		// Marshal provided interface into JSON structure
// 		uj, _ := json.Marshal(u)

// 		// Write content-type, statuscode, payload
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(201)
// 		fmt.Fprintf(w, "%s", uj)
// 	})

// 	r.DELETE("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		// TODO: only write status for now
// 		w.WriteHeader(200)
// 	})

// 	http.ListenAndServe("localhost:3000", r)
// }
