package users

import (
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

func (h *usersHandler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

func (h *usersHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
