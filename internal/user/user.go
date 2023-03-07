package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}

type userServiceInterface interface {
	Register(User) error
}

type userRepositoryInterface interface {
	Create(User) error
}
