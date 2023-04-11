package article

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tigor7/go-chi-realworld-example-app/internal/auth"
)

type articleHandler struct {
	articleService articleServiceInterface
}

func NewarticleHandler(s articleServiceInterface) articleHandler {
	return articleHandler{
		articleService: s,
	}
}

func (h *articleHandler) RegisterRoutes(r *chi.Mux) {
	// Auth routes
	r.Group(func(r chi.Router) {
		r.Use(auth.ValidateJWT)
		r.Post("/api/aticles", h.handleCreateArticle)
	})
}

type createArticleRequest struct {
	Article struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"article"`
}

func (h *articleHandler) handleCreateArticle(w http.ResponseWriter, r *http.Request) {

}
