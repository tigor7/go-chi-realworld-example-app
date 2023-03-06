package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type userHandler struct {
	userService userServiceInterface
}

func NewUserHandler(s userServiceInterface) userHandler {
	return userHandler{
		userService: s,
	}
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
	u := User{
		Username: request.User.Username,
		Email:    request.User.Email,
		Password: request.User.Password,
	}
	if err := h.userService.Register(u); err != nil {

	}
	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
