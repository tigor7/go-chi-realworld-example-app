package user

import (
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
