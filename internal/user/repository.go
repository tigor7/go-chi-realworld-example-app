package user

import (
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
