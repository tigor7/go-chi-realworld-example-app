package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/tigor7/go-chi-realworld-example-app/internal/httputil"
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
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Password string  `json:"password"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
	} `json:"user"`
}

func (r *registerRequest) Validate() error {
	return validation.ValidateStruct(&r.User,
		validation.Field(&r.User.Username, validation.Required, validation.Length(2, 255)),
		validation.Field(&r.User.Email, validation.Required, validation.Length(0, 255), is.Email),
		validation.Field(&r.User.Password, validation.Required, validation.Length(8, 255)),
	)
}

func (h *userHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	request := registerRequest{}
	if err := httputil.BindAndValidate(r, &request); err != nil {
		httputil.RespondErr(w, http.StatusUnprocessableEntity, err)
		return
	}
	u := User{
		Username: request.User.Username,
		Email:    request.User.Email,
		Password: request.User.Password,
		Bio:      request.User.Bio,
		Image:    request.User.Image,
	}
	token, err := h.userService.Register(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httputil.Respond(w, http.StatusOK, NewUserResponse(u, token))
}

func (h *userHandler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
