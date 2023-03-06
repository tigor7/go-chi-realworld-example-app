package user

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}

type userServiceInterface interface {
}

type userRepositoryInterface interface {
	Create(User) error
}
