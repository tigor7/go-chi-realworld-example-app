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

func (r *userRepository) Create(u User) error {
	_, err := r.db.NamedExec("INSERT INTO users (id, username, email, password) VALUES (:id, :username, :email, :password)", u)
	return err
}
