package article

import "github.com/google/uuid"

type Article struct {
	ID          uuid.UUID
	Title       string
	Description string
	Body        string
	TagList     *[]string
}

type articleRepositoryInterface interface {
}

type articleServiceInterface interface {
}
