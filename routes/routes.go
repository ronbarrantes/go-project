package routes

import "net/http"

type UserHandler struct{}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request)   {}
func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request)   {}
func (u *UserHandler) Delete(w http.ResponseWriter, r *http.Request)   {}
func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {}
func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request)  {}

func New() *UserHandler {
	return &UserHandler{}
}
