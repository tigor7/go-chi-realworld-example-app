package user

import "github.com/jmoiron/sqlx"

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) userRepositoryInterface {
	return &userRepository{
		db: db,
	}
}
