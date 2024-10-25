package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"ronb.co/project/utils"
)

var users = make(map[string]User)

// ##### TYPES #####
type APIServer struct {
	listenAddress string
}

type APIError struct {
	Error string
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Store struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// ##### SERVER #####
func Server(listerAddr string) *APIServer {
	return &APIServer{
		listenAddress: listerAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", s.handleAccount)
	router.HandleFunc("/api/users/{id}", s.handleGetUser)

	fmt.Printf("Listening on %s", s.listenAddress)
	log.Fatal(http.ListenAndServe(":"+s.listenAddress, router))
}

// ##### USER MANAGEMENT #####

func generateUser(firstName, lastName string) (User, error) {
	randId, err := utils.GenerateRandomId()
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:        randId,
		FirstName: firstName,
		LastName:  lastName,
	}

	return user, nil
}

// GetUsers
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleGetUsers(w, r)
	case "POST":
		s.handleCreateUser(w, r)
	default:
		http.Error(w, fmt.Sprintf("Method not allowed: %s", r.Method), http.StatusMethodNotAllowed)
	}
}

// CREATE USER
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := User{}

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if newUser.FirstName == "" || newUser.LastName == "" {
		http.Error(w, "firstName and lastName are required fields", http.StatusBadRequest)
		return
	}

	createdUser, err := generateUser(newUser.FirstName, newUser.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Append the new user to the users slice
	// users = append(users, createdUser)
	users[createdUser.ID] = createdUser

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// READ USER
func (s *APIServer) handleGetUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	arrayUsers, err := utils.MakeMapToArray(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(arrayUsers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	fmt.Println(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users[id]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
