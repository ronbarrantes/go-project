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
	OwnerId   string `json:"ownerId"`
	StoreName string `json:"storeName"`
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
	router.HandleFunc("/api/users/{id}", s.handleSingleUser)

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

func (s *APIServer) handleSingleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleGetUserById(w, r)
	case "PUT":
		s.handleUpdateUserById(w, r)
	case "DELETE":
		s.handleDeleteUserById(w, r)
	default:
		http.Error(w, fmt.Sprintf("Method not allowed: %s", r.Method), http.StatusMethodNotAllowed)
	}

}

// CreateUser
func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// Decode the JSON body into the struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if user.FirstName == "" || user.LastName == "" {
		http.Error(w, "firstName and lastName are required fields", http.StatusBadRequest)
		return
	}

	createdUser, err := generateUser(user.FirstName, user.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users[createdUser.ID] = createdUser
	utils.WriteJSON(w, http.StatusCreated, createdUser)
}

// READ USER
func (s *APIServer) handleGetUsers(w http.ResponseWriter, _ *http.Request) {
	arrayUsers, err := utils.MakeMapToArray(users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	utils.WriteJSON(w, http.StatusOK, arrayUsers)
}

// GetUserById
func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// Check if the user exists
	response := map[string]string{"message": "User not found"}
	if _, exists := users[id]; !exists {
		utils.WriteJSON(w, http.StatusNotFound, response)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users[id])
}

func (s *APIServer) handleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// Check if the user exists

	existingUser, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the JSON body into a map
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Apply the updates to the existing user
	if firstName, ok := updates["firstName"].(string); ok {
		existingUser.FirstName = firstName
	}

	if lastName, ok := updates["lastName"].(string); ok {
		existingUser.LastName = lastName
	}

	users[id] = existingUser
	utils.WriteJSON(w, http.StatusOK, users[id])
}

func (s *APIServer) handleDeleteUserById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	// Check if the user exists
	response := map[string]string{"message": "User not found"}
	if _, exists := users[id]; !exists {
		utils.WriteJSON(w, http.StatusNotFound, response)
		return
	}

	// Delete the user
	delete(users, id)
	// Create a response message
	response = map[string]string{"message": "Item deleted"}
	utils.WriteJSON(w, http.StatusOK, response)
}

// STORE STUFF
func generateStore(ownerId, storeName string) (Store, error) {
	randId, err := utils.GenerateRandomId()
	if err != nil {
		return Store{}, err
	}

	store := Store{
		ID:        randId,
		StoreName: storeName,
		OwnerId:   ownerId,
	}

	return store, nil
}

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
