package users

import (
	"encoding/json"
	"fmt"
	"github.com/Modular-Company/modular-bar/internal/domain/users"
	"net/http"
)

type UsersHandler struct {
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

func (h *UsersHandler) ListUsers(w http.ResponseWriter, r *http.Request) {

}

func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, "cannot parse request payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("user received: %v", userDTO)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]users.User{"user": *userDTO.ToDomain()})
}

func (h *UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}
