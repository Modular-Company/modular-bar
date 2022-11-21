package entrypoint

import (
	"github.com/Modular-Company/modular-bar/internal/infra/entrypoint/users"
	"net/http"

	"github.com/gorilla/mux"
)

func Load(usersHandler *users.UsersHandler) (*mux.Router, error) {
	router := mux.NewRouter()

	router.HandleFunc("/users", usersHandler.ListUsers).Methods(http.MethodGet)
	router.HandleFunc("/users", usersHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{user_id}", usersHandler.GetUser).Methods(http.MethodGet)

	return router, nil
}
