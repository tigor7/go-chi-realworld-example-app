package user

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type userServiceInterface interface {
}

type userRepositoryInterface interface {
}
