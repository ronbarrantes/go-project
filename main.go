package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := "8000"

	http.HandleFunc("/users", GetUsers)

	fmt.Printf("Listening on %s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

// make the routes

// USER STUFF

type User struct {
  ID string `json:"id"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
}

func MakeUser (firstName , lastName string) *User{
  user := User{
    ID: ,
   FirstName: firstName, 
    LastName: lastName,
  }
  return &user 
}

// GetUsers
func GetUsers(w http.ResponseWriter, r *http.Request) {
  
  users := []User{

  } 


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
