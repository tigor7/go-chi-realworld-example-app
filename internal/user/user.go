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
}

type userRepositoryInterface interface {
	Create(User) error
}
