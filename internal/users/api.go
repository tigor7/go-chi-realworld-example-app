package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type usersHandler struct {
}

func NewUsersHandler() usersHandler {
	return usersHandler{}
}

func (h *usersHandler) RegisterRoutes(r *chi.Mux) {
	r.Post("/api/users", h.handleRegister)
	r.Post("/api/users/login", h.handleLogin)
}

type registerRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (h *usersHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	request := registerRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}

func (h *usersHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
