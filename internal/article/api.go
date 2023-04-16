package article

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/tigor7/go-chi-realworld-example-app/internal/auth"
	"github.com/tigor7/go-chi-realworld-example-app/internal/httputil"
)

type articleHandler struct {
	articleService articleServiceInterface
}

func NewArticleHandler(s articleServiceInterface) articleHandler {
	return articleHandler{
		articleService: s,
	}
}

func (h *articleHandler) RegisterRoutes(r *chi.Mux) {
	// Auth routes

	r.Group(func(r chi.Router) {
		r.Use(auth.ValidateJWT)
		r.Post("/api/articles", h.handleCreateArticle)
	})
}

type createArticleRequest struct {
	Article struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"article"`
}

func (r createArticleRequest) Validate() error {
	return validation.ValidateStruct(&r.Article,
		validation.Field(&r.Article.Title, validation.Required),
		validation.Field(&r.Article.Description, validation.Required),
		validation.Field(&r.Article.Body, validation.Required),
	)
}
func (h *articleHandler) handleCreateArticle(w http.ResponseWriter, r *http.Request) {
	request := createArticleRequest{}
	if err := httputil.BindAndValidate(r, &request); err != nil {
		httputil.RespondErr(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid := uidFromRequest(r)
	article := Article{
		Title:       request.Article.Title,
		Description: request.Article.Description,
		Body:        request.Article.Body,
		AuthorID:    uid,
	}
	if err := h.articleService.Create(article); err != nil {
		httputil.RespondErr(w, http.StatusUnprocessableEntity, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func uidFromRequest(r *http.Request) uuid.UUID {
	uid, _ := uuid.Parse(r.Context().Value("uid").(string))
	return uid
}
