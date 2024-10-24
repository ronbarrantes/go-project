package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	"ronb.co/project/utils"
)

func main() {
	PORT := "8000"

	// router := mux.NewRouter()

	http.HandleFunc("/users", GetUsers)
	fmt.Printf("Listening on %s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

// make the routes

// USER STUFF

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func MakeUser(firstName, lastName string) *User {
	user := User{
		ID:        utils.GenerateRandomId(),
		FirstName: firstName,
		LastName:  lastName,
	}
	return &user
}

// GetUsers
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		*MakeUser("Ron", "B"),
		*MakeUser("Jen", "M"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

type APIError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// GetUserById

// MakeUser
// UpdateUser
// DeleteUser

// STORE STUFF
// CreateStores
// GetStoreById
// GetStores
// UpdateStore
// DeleteStore

// AUTHENTICATION STUFF
// IsAuthenticated
// LogIn
// LogOut
// SignUp

// DATABASE STUFF
