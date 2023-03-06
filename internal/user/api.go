package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type userHandler struct {
}

func NewUserHandler() userHandler {
	return userHandler{}
}

func (h *userHandler) RegisterRoutes(r *chi.Mux) {
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

func (h *userHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	request := registerRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
}

func (h *userHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
