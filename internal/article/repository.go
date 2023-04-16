package article

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type articleRespository struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) articleRepositoryInterface {
	return &articleRespository{
		db: db,
	}
}

func (r articleRespository) CreateArticle(a Article) error {
	a.ID = uuid.New()
	_, err := r.db.NamedExec("INSERT INTO articles (id, title, slug, description, body, author_id) VALUES (:id, :title, :slug, :description, :body, :author_id)", a)
	return err
}
