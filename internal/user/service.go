package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tigor7/go-chi-realworld-example-app/internal/auth"
)

type userService struct {
	userRepository userRepositoryInterface
}

func NewUserService(r userRepositoryInterface) userServiceInterface {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) Register(u User) (string, error) {
	u.ID = uuid.New()
	err := s.userRepository.Create(u)
	if err != nil {
		return "", err
	}

	token, err := auth.CreateJWT(u.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) Login(u User) (User, string, error) {
	us, err := s.userRepository.GetByEmail(u.Email)
	if err != nil {
		return us, "", err
	}
	if !ComparePassword(us.Password, u.Password) {
		return us, "", errors.New("Username and password do not match")
	}
	token, err := auth.CreateJWT(us.ID)
	if err != nil {
		return us, "", err
	}
	return us, token, nil
}
