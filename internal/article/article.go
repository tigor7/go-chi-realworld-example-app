package article

import "github.com/google/uuid"

type Article struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Slug        string    `db:"slug"`
	Body        string    `db:"body"`
	TagList     *[]string
	AuthorID    uuid.UUID `db:"author_id"`
}

type articleRepositoryInterface interface {
	CreateArticle(a Article) error
}

type articleServiceInterface interface {
	Create(a Article) error
}
