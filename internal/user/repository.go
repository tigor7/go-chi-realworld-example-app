package user

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) userRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByEmail(email string) (User, error) {
	u := User{}
	err := r.db.Get(&u, "SELECT * FROM users WHERE email=$1", email)
	if err == sql.ErrNoRows {
		return u, errors.New("User not found")
	}
	return u, err
}

func (r *userRepository) GetByUsername(username string) (User, error) {
	u := User{}
	err := r.db.Get(&u, "SELECT * FROM users WHERE username=$1", username)
	if err == sql.ErrNoRows {
		return u, errors.New("User not found")
	}
	return u, err
}

func (r *userRepository) GetByUserID(id uuid.UUID) (User, error) {
	u := User{}
	err := r.db.Get(&u, "SELECT * FROM users WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return u, errors.New("User not found")
	}
	return u, err
}

func (r *userRepository) Create(u User) error {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	_, err = r.db.NamedExec("INSERT INTO users (id, username, email, password, bio, image) VALUES (:id, :username, :email, :password, :bio, :image)", u)
	return err
}

func (r *userRepository) Follow(uid uuid.UUID, friendID uuid.UUID) error {
	_, err := r.db.Exec("INSERT INTO follows (user_id, followed_id) SELECT $1, $2 WHERE NOT EXISTS(SELECT * FROM follows WHERE user_id=$1 AND followed_id=$2)", uid, friendID)
	return err
}

func (r *userRepository) Unfollow(uid uuid.UUID, friendID uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM follows WHERE user_id=$1 AND followed_id=$2", uid, friendID)
	return err
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
