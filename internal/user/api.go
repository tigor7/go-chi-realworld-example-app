package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/tigor7/go-chi-realworld-example-app/internal/auth"
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
	r.Get("/api/profiles/{username}", h.handleGetProfile)

	// Auth routes
	r.Group(func(r chi.Router) {
		r.Use(auth.ValidateJWT)
		r.Get("/api/user", h.handleGetUser)
		r.Post("/api/profiles/{username}/follow", h.handleFollow)
		r.Delete("/api/profiles/{username}/follow", h.handleUnfollow)

	})

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

type loginRequest struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (r *loginRequest) Validate() error {
	return validation.ValidateStruct(&r.User,
		validation.Field(&r.User.Email, validation.Required, validation.Length(0, 255), is.Email),
		validation.Field(&r.User.Password, validation.Required, validation.Length(8, 255)),
	)
}

func (h *userHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	request := loginRequest{}
	if err := httputil.BindAndValidate(r, &request); err != nil {
		httputil.RespondErr(w, http.StatusUnprocessableEntity, err)
		return
	}
	u := User{
		Email:    request.User.Email,
		Password: request.User.Password,
	}
	us, token, err := h.userService.Login(u)
	if err != nil {
		httputil.RespondErr(w, http.StatusInternalServerError, err)
		return
	}
	httputil.Respond(w, http.StatusOK, NewUserResponse(us, token))
}

func (h *userHandler) handleGetProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	u, err := h.userService.GetProfile(username)
	if err != nil {
		httputil.RespondErr(w, http.StatusInternalServerError, err)
		return
	}
	httputil.Respond(w, http.StatusOK, NewProfileResponse(u))
}

func (h *userHandler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	uid := uidFromRequest(r)
	u, err := h.userService.GetUserByID(uid)
	if err != nil {
		httputil.RespondErr(w, http.StatusInternalServerError, err)
		return
	}
	httputil.Respond(w, http.StatusOK, NewUserResponse(u, ""))
}

func (h *userHandler) handleFollow(w http.ResponseWriter, r *http.Request) {
	uid := uidFromRequest(r)
	username := chi.URLParam(r, "username")

	friend, err := h.userService.Follow(uid, username)
	if err != nil {
		httputil.RespondErr(w, http.StatusInternalServerError, err)
		return
	}
	httputil.Respond(w, http.StatusOK, NewProfileResponse(friend))
}

func (h *userHandler) handleUnfollow(w http.ResponseWriter, r *http.Request) {
	uid := uidFromRequest(r)
	username := chi.URLParam(r, "username")

	friend, err := h.userService.Unfollow(uid, username)
	if err != nil {
		httputil.RespondErr(w, http.StatusInternalServerError, err)
		return
	}
	httputil.Respond(w, http.StatusOK, NewProfileResponse(friend))
}

func uidFromRequest(r *http.Request) uuid.UUID {
	uid, _ := uuid.Parse(r.Context().Value("uid").(string))
	return uid
}
