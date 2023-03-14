package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	Bio      *string   `db:"bio"`
	Image    *string   `db:"image"`
	Password string    `db:"password"`
}

type userServiceInterface interface {
	Register(User) (token string, err error)
	Login(User) (User, string, error)

	GetProfile(username string) (User, error)
}

type userRepositoryInterface interface {
	GetByEmail(email string) (User, error)
	GetByUsername(username string) (User, error)
	Create(User) error
}
