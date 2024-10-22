package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// HANDLE THE ERROR
			WriteJSON(w, http.StatusBadRequest, APIError{
				Error: err.Error(),
			})
		}
	}
}

type APIServer struct {
	listerAddr string
}

func NewAPIServer(listerAddr string) *APIServer {
	return &APIServer{
		listerAddr: listerAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	// router.HandleFunc("/accout", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API Running on Port: ", s.listerAddr)
	log.Fatal(http.ListenAndServe(s.listerAddr, router))

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)

	case "POST":
		return s.handleCreateAccount(w, r)

	case "DELETE":
		return s.handleDeleteAccount(w, r)

	default:
		return fmt.Errorf("Method not allowed %s", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
