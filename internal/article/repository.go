package article

import "github.com/jmoiron/sqlx"

type articleRespository struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) articleRepositoryInterface {
	return &articleRespository{
		db: db,
	}
}
